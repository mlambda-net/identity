package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/mlambda-net/identity/pkg/domain/entity"
	"github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/spec"
  "github.com/mlambda-net/identity/pkg/domain/utils"
	"github.com/mlambda-net/monads/monad"
)

type identityStore struct {
	db *pg.DB
}

func (i *identityStore) ByEmail(email string) monad.Mono {
	 user := &entity.Identity{}
	_, err := i.db.QueryOne(user, `SELECT * FROM identities Where email = ?`, email)
	if err != nil {
		return monad.ToMono(err)
	}

	return monad.ToMono(user)
}

func (i *identityStore) ById(id int64) monad.Mono {
	var user entity.Identity
	_, err := i.db.QueryOne(&user, `SELECT * FROM identities Where id = ?`, id)
	if err != nil {
		return monad.ToMono(err)
	}

	return monad.ToMono(&user)
}

func (i *identityStore) Update(user *entity.Identity) monad.Mono {
	_, err := i.db.Model(user).WherePK().Update()
	if err != nil {
		return monad.ToMono(err)
	}
	return monad.ToMono(user)
}

func (i *identityStore) Save(id *entity.Identity) monad.Mono {
	_, err := i.db.Model(id).Insert()
	if err != nil {
		return monad.ToMono(err)
	}
	return monad.ToMono(id)
}

func (i *identityStore) Delete(id int64) monad.Mono {
	_, err := i.db.Model(&entity.Identity{
		Id: id,
	}).WherePK().Delete()

	if err != nil {
		return monad.ToMono(err)
	}
	return monad.ToMono(nil)
}

func (i *identityStore) All(spec spec.Spec) monad.Mono {
  var items []entity.Identity

  _, err := i.db.Query(&items, fmt.Sprintf("SELECT * FROM identities where %s", spec.Query()))

  if err != nil {
    monad.ToMono(err)
  }

  return monad.ToMono(items)
}

func (i *identityStore) Single(spec spec.Spec) monad.Mono {
  var items []entity.Identity
  _, err := i.db.Query(&items, fmt.Sprintf("SELECT * FROM identities where %s", spec.Query()))

  if err != nil {
    monad.ToMono(err)
  }

  if len(items) > 0 {
    return monad.ToMono(items[0])
  }

  return monad.ToMono(nil)
}


func NewIdentityStore(config *utils.Configuration) repository.IdentityStore {

	db := pg.Connect(&pg.Options{
		User:     config.Db.User,
		Password: config.Db.Password,
		Addr:     fmt.Sprintf("%s:%s", config.Db.Host, config.Db.Port),
		Database: config.Db.Schema,
	})

	return &identityStore{db: db}

}

func (i *identityStore) Close() {
	_ = i.db.Close()
}
