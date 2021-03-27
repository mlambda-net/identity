package services

import (
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/specs"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/mlambda-net/identity/pkg/infrastructure/db"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
)

type AppQuery interface {
  Search(filter string) monad.Mono
  Find(id string) monad.Mono
}

type appQuery struct {
  repo repository.AppRepository
}

func (a appQuery) Find(id string) monad.Mono {
  return a.repo.Get(spec.Specify(specs.ById(id)))
}

func (a appQuery) Search(filter string) monad.Mono {
  return a.repo.Search(spec.Specify(spec.Empty()))
}

func NewAppQuery(config *utils.Configuration) AppQuery {
  repo := db.NewAppRepository(config)
  return &appQuery{repo: repo}
}
