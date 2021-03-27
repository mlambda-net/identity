package db

import (
  "github.com/go-pg/pg/v10"
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
)

type roleRepository struct {
  db *pg.DB
}

func (r roleRepository) Get(filter spec.Spec) monad.Mono {
  return r.Search(filter).Bind(func(any types.Any) monad.Mono {
    items := any.([]*entity.Role)
    if len(items) == 0 {
      return monad.ToMono(monad.Empty())
    }
    return monad.ToMono(monad.Unit(items[0]))
  })
}

func (r roleRepository) Search(filter spec.Spec) monad.Mono {
  var items []*entity.Role

  query := filter.Query(`SELECT r.id, r.name, r.description, a.id as app__id, a.name as app__name
  FROM roles as r
  INNER JOIN app as a ON r.app = a.id`)

  _, err := r.db.Query(&items, query, filter.Data()...)
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(items)
}

func (r roleRepository) Add(role *entity.Role) monad.Mono {
  _, err := r.db.Exec("call role_add(?, ?, ?, ?)", role.ID, role.App.ID, role.Name, role.Description)
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(role)
}

func (r roleRepository) Update(role *entity.Role) monad.Mono {
  _, err := r.db.Exec("call role_update(?, ?)", role.ID, role.Description)
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(role)
}

func (r roleRepository) Delete(id uuid.UUID) monad.Mono {
  _, err := r.db.Exec("call role_delete(?)", id)
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(id)
}

func NewRoleRepository(config *utils.Configuration) repository.RoleRepository {
  db := initializeDB(config)
  return &roleRepository{db : db}
}
