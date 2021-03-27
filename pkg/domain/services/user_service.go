package services

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/specs"
  utils "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/mlambda-net/identity/pkg/infrastructure/db"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/ex"
  "github.com/mlambda-net/net/pkg/spec"
)

type UserService interface {
	Create(user *entity.Identity) monad.Mono
	Delete(id uuid.UUID) monad.Mono
	Update(user *entity.Identity) monad.Mono
	ChangePassword(email string, password string) monad.Mono
	Authenticate(login string, password string) monad.Mono
}

type userService struct {
	repo  repository.IdentityStore
}

func (s userService) Authenticate(login string, password string) monad.Mono {
  return s.repo.Single(spec.Specify(specs.ByEmail(login))).Bind(func(any types.Any) monad.Mono {
    usr := any.(*entity.Identity)
    actual, e:= utils.Decrypt(usr.Password)
    if e != nil {
      return monad.ToMono(ex.Error(e))
    }
    return monad.ToMono(actual == password)
  })
}


func (s userService) ChangePassword(email string, password string) monad.Mono {
  newPass, err := utils.Encrypt(password)
  if err != nil {
    return monad.ToMono(ex.Error(err))
  }

  return s.repo.Single(spec.Specify(specs.ByEmail(email))).Bind(func(any types.Any) monad.Mono {
    user := any.(*entity.Identity)
    user.Password = newPass
    return s.repo.Update(user)
  })
}

func (s userService) Create(user *entity.Identity) monad.Mono {
  newPass, err := utils.Encrypt(user.Password)
  if err != nil {
    monad.ToMono(ex.Error(err))
  }
  user.Password = newPass

	return s.repo.Save(user)

}

func (s userService) Delete(id uuid.UUID) monad.Mono {
	return s.repo.Delete(id)
}

func (s userService) Update(user *entity.Identity) monad.Mono {
	return s.repo.Update(user)
}

func NewUserService(config *utils.Configuration) UserService {
  repo := db.NewIdentityStore(config)
  return userService{
    repo:  repo,
  }
}
