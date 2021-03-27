package repository

import (
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
)

type AppRepository interface {
  Save(entity *entity.App) monad.Mono
	Search(filter spec.Spec) monad.Mono
  Get(filter spec.Spec) monad.Mono
	Update(entity *entity.App) monad.Mono
}
