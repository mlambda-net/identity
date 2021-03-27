package repository

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
)

type RoleRepository interface {
  Get(filter spec.Spec) monad.Mono
  Search(filter spec.Spec) monad.Mono
  Add(rol *entity.Role) monad.Mono
  Update(rol *entity.Role) monad.Mono
  Delete(id uuid.UUID) monad.Mono
}
