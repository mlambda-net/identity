package services

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/specs"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/mlambda-net/identity/pkg/infrastructure/db"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
)

type RolesQuery interface {
  Find(id uuid.UUID) monad.Mono
  Search(filter string) monad.Mono
}

type rolesQuery struct {
  repo repository.RoleRepository
}

func (r rolesQuery) Find(id uuid.UUID) monad.Mono {
  return r.repo.Get(spec.Specify(specs.ByRoleId(id.String())))
}

func (r rolesQuery) Search(filter string) monad.Mono {
  return r.repo.Search(spec.Specify(specs.ByRole(filter)))
}

func NewRolesQuery(config *utils.Configuration) RolesQuery {
  repo := db.NewRoleRepository(config)
  return &rolesQuery{repo: repo}
}
