package services

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/mlambda-net/identity/pkg/infrastructure/db"
  "github.com/mlambda-net/monads/monad"
)

type RolesService interface {
  Add(role *entity.Role) monad.Mono
  Update(role *entity.Role) monad.Mono
  Delete(id uuid.UUID) monad.Mono
}

type rolesService struct {
  repo repository.RoleRepository
}

func (r rolesService) Add(role *entity.Role) monad.Mono {
  return r.repo.Add(role)
}

func (r rolesService) Update(role *entity.Role) monad.Mono {
  return r.repo.Update(role)
}

func (r rolesService) Delete(id uuid.UUID) monad.Mono {
  return r.repo.Delete(id)
}

func NewRolesService(config *utils.Configuration) RolesService {
  repo := db.NewRoleRepository(config)
  return &rolesService{ repo : repo }
}
