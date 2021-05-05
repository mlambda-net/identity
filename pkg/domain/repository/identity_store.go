package repository

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
)

type IdentityStore interface {
	Save(id *entity.Identity) monad.Mono
	Delete(id uuid.UUID) monad.Mono
	Update(user *entity.Identity) monad.Mono
	Single(spec spec.Spec) monad.Mono
	All(spec spec.Spec) monad.Mono
	Rights(id uuid.UUID) monad.Mono
	Close()
}
