package services

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/specs"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/mlambda-net/identity/pkg/infrastructure/db"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
)

type QueryService interface {
  Query(query spec.Expression) monad.Mono
  Single(id uuid.UUID) monad.Mono
  ByEmail(email string) monad.Mono
	All(filter string) monad.Mono
}

type queryService struct {
  repo repository.IdentityStore
}

func (s *queryService) All(filter string) monad.Mono {
  return s.repo.All(spec.Specify(specs.ByUser(filter)))
}

func (s *queryService) ByEmail(email string) monad.Mono {
  return s.repo.Single(spec.Specify(specs.ByEmail(email))).Bind(func(any types.Any) monad.Mono {
    user := any.(*entity.Identity)
    return s.repo.Rights(user.ID).Bind(func(any types.Any) monad.Mono {
      user.Roles = any.([]*entity.Role)
      return monad.ToMono(user)
    })
  })
}

func (s *queryService) Single(id uuid.UUID) monad.Mono {
  return s.repo.Single(spec.Specify(specs.ById(id.String()))).Bind(func(any types.Any) monad.Mono {
    user := any.(*entity.Identity)
    return s.repo.Rights(user.ID).Bind(func(any types.Any) monad.Mono {
      user.Roles = any.([]*entity.Role)
      return monad.ToMono(user)
    })
  })
}

func (s *queryService) Query(query spec.Expression) monad.Mono {
  return s.repo.All(spec.Specify(query))
}

func NewQueryService(config *utils.Configuration) QueryService {
  repo := db.NewIdentityStore(config)

  return &queryService{
    repo:  repo,

  }
}
