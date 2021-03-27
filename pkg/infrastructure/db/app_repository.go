package db

import (
  "github.com/go-pg/pg/v10"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
)

type appRepository struct {
  db *pg.DB
}

func (a appRepository) Update(entity *entity.App) monad.Mono {
  _, err := a.db.Exec("call app_update(?, ?)",entity.ID, entity.Description)
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(entity)
}

func (a appRepository) Get(filter spec.Spec) monad.Mono {
  return a.Search(filter).Bind(func(any types.Any) monad.Mono {
    items := any.([]*entity.App)
    if len(items) == 0 {
      return monad.ToMono(monad.Empty())
    }
    return monad.ToMono(monad.Unit(items[0]))
  })
}

func (a appRepository) Search(filter spec.Spec) monad.Mono {
  var items []*entity.App
  var err error
  _, err = a.db.Query(&items,  filter.Query("SELECT id, name, description FROM app"), filter.Data()...)
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(items)
}

func (a appRepository) Save(entity *entity.App) monad.Mono {
  _, err := a.db.Exec("call app_add(?, ?, ?)",entity.ID, entity.Name, entity.Description)
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(entity)
}

func NewAppRepository(config *utils.Configuration) repository.AppRepository {
  db := initializeDB(config)
  return &appRepository{db : db}
}
