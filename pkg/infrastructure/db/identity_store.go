package db

import (
  "github.com/go-pg/pg/v10"
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/ex"
  "github.com/mlambda-net/net/pkg/spec"
)

type identityStore struct {
	db *pg.DB
}

func (i *identityStore) Update(user *entity.Identity) monad.Mono {
  begin, err := i.db.Begin()
  if err != nil {
    return monad.ToMono(err)
  }

  err = i.purgeRights(user, begin)
  if err != nil {
    return monad.ToMono(err)
  }

  _, err = i.db.Exec("call user_update(?,?,?)", user.ID, user.Name, user.LastName)
  if err != nil {
    err := begin.Rollback()
    if err != nil {
      return monad.ToMono(err)
    }
    return monad.ToMono(err)
  }

  err  = i.addRights(user, begin)
  if err != nil {
    return monad.ToMono(err)
  }

  err = begin.Close()
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(user)
}

func (i *identityStore) Save(user *entity.Identity) monad.Mono {
  begin, err := i.db.Begin()
  if err != nil {
    return monad.ToMono(err)
  }
  err = i.purgeRights(user, begin)
  if err != nil {
    return monad.ToMono(err)
  }

  _, err = i.db.Exec("call user_add(?,?,?,?,?)", user.ID, user.Name, user.LastName, user.Email, user.Password)
  if err != nil {
    e := begin.Rollback()
    if e != nil {
      return monad.ToMono(e)
    }

    return monad.ToMono(ex.Friend("the user can not be saved", err))
  }

  err = i.addRights(user, begin)
  if err != nil {
    return monad.ToMono(err)
  }

  err = begin.Close()
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(user)
}

func (i *identityStore) Delete(id uuid.UUID) monad.Mono {
	_, err := i.db.Model(&entity.Identity{
		ID: id,
	}).WherePK().Delete()

	if err != nil {
		return monad.ToMono(err)
	}
	return monad.ToMono(nil)
}

func (i *identityStore) All(spec spec.Spec) monad.Mono {var items []*entity.Identity
  var err error
  if spec != nil {
    query := spec.Query("SELECT id, name, last_name, email FROM identities")
    data := spec.Data()
    _, err = i.db.Query(&items, query , data...)
   } else {
    _, err = i.db.Query(&items, "SELECT id, name, last_name, email FROM identities")
  }

  if err != nil {
    monad.ToMono(err)
  }
  return monad.ToMono(items)
}

func (i *identityStore) Single(spec spec.Spec) monad.Mono {
  var items []*entity.Identity

  query := spec.Query( "SELECT * FROM identities")
  _, err := i.db.Query(&items, query, spec.Data()...)

  if err != nil {
    monad.ToMono(err)
  }

  if len(items) > 0 {
    return monad.ToMono(items[0])
  }

  return monad.ToMono(nil)
}

func (i *identityStore) Rights( id uuid.UUID) monad.Mono {
  var items []*entity.Role
  query := `SELECT ro.id, ro.name, ro.description, a.id as app__id, a.name as app__name
  FROM Roles  ro
  INNER JOIN rights ri
  ON ro.id = ri.roleID
  INNER JOIN App a
  ON ro.app = a.id
  WHERE ri.userid = ?`

  _, err := i.db.Query(&items,query, id)

  if err != nil {
    return monad.ToMono(err)
  }
  return  monad.ToMono(items)
}

func (i *identityStore) addRights(user *entity.Identity, begin *pg.Tx) error {
  for _, rol := range user.Roles {
    _, err := i.db.Exec("call rights_add(?,?)", user.ID, rol.ID)
    if err != nil {
      err := begin.Rollback()
      if err != nil {
        return err
      }
      return err
    }
  }
  return nil
}

func (i *identityStore) purgeRights(user *entity.Identity, begin *pg.Tx) error {
  _, err := i.db.Exec("call rights_purge(?)", user.ID)
  if err != nil {
    err := begin.Rollback()
    if err != nil {
      return err
    }
    return err
  }
  return nil
}

func NewIdentityStore(config *utils.Configuration) repository.IdentityStore {
	db := initializeDB(config)
	return &identityStore{db: db}
}

func (i *identityStore) Close() {
	_ = i.db.Close()
}
