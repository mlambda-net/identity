package services

import (
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/spec"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/mlambda-net/identity/pkg/infrastructure/db"
)

type QueryService interface {
  Query(query spec.Expression) ([]entity.Identity, error)
  Single(id int64) (*entity.Identity, error)
  ByEmail(email string) (*entity.Identity, error)
}

type queryService struct {
  repo repository.IdentityStore
}

func (s *queryService) ByEmail(email string) (*entity.Identity, error) {
  rsp, err := s.repo.ByEmail(email).Unwrap()
  if err != nil {
    return nil, err
  }

  user := rsp.(*entity.Identity)
  return user, nil
}

func (s *queryService) Single(id int64) (*entity.Identity, error) {
  u, e := s.repo.ById(id).Unwrap()
  if e != nil {
    return nil, e
  }
  return u.(*entity.Identity), nil
}

func (s *queryService) Query(query spec.Expression) ([]entity.Identity, error) {
  r := s.repo.All(spec.Specify(query))
  data, err := r.Unwrap()
  if err != nil {
    return nil, err
  }

  return data.([]entity.Identity), nil
}

func NewQueryService(config *utils.Configuration) QueryService {
  repo := db.NewIdentityStore(config)

  return &queryService{
    repo:  repo,

  }
}
