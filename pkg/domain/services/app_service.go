package services

import (
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/mlambda-net/identity/pkg/infrastructure/db"
  "github.com/mlambda-net/monads/monad"
)

type AppService interface {
  Add(entity *entity.App) monad.Mono
	Update(entity *entity.App) monad.Mono
}

type appService struct {
  repo repository.AppRepository
}

func (a *appService) Update(entity *entity.App) monad.Mono {
  return a.repo.Update(entity)
}

func (a *appService) Add(entity *entity.App) monad.Mono {
  return a.repo.Save(entity)
}

func NewAppService(config *utils.Configuration) AppService{
  repo := db.NewAppRepository(config)
  return &appService{ repo: repo }
}
