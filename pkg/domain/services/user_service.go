package services

import (
	"errors"
	"fmt"
	"github.com/mlambda-net/identity/pkg/domain/entity"
	"github.com/mlambda-net/identity/pkg/domain/repository"
	"github.com/mlambda-net/identity/pkg/domain/spec"
	"github.com/mlambda-net/identity/pkg/domain/utils"
	"github.com/mlambda-net/identity/pkg/infrastructure/endpoint/db"
	"github.com/mlambda-net/net/pkg/security"
	"os"
)

type UserService interface {
	Create(user *entity.Identity) (*entity.Identity, error)
	Query(query spec.Expression) ([]entity.Identity, error)
	Delete(id int64) error
	Update(user *entity.Identity) error
	Single(id int64) (*entity.Identity, error)
	ChangePassword(email string, password string) error
	ByEmail(email string) (*entity.Identity, error)
    Authenticate(login string, password string) (string, error)
}

type userService struct {
	repo  repository.IdentityStore
	query repository.IdentityQuery
}

func (s userService) Authenticate(login string, password string) (string, error) {
	token := security.NewToken(os.Getenv("SECRET_KEY"))
	rsp, err := s.repo.ByEmail(login).Unwrap()
	if err != nil {
		return "", err
	}
	usr := rsp.(*entity.Identity)

	systemPwd, err := utils.Decrypt(usr.Password)
	if err != nil {
		return "", err
	}

	if systemPwd != password {
		return "", errors.New("user or password mismatch")
	}

	return token.Create(map[string]interface{}{
		"authorize": true,
		"id":         usr.Id,
		"name":       fmt.Sprintf("%s %s", usr.Name, usr.LastName),
		"email":      usr.Email,
	})
}

func (s userService) ByEmail(email string) (*entity.Identity, error) {
  rsp, err := s.repo.ByEmail(email).Unwrap()
  if err != nil {
    return nil, err
  }

  user := rsp.(*entity.Identity)
  return user, nil
}

func (s userService) ChangePassword(email string, password string) error {
	newPass, err := utils.Encrypt(password)
	if err != nil {
		return err
	}

	 rsp, err := s.repo.ByEmail(email).Unwrap()
	 if err != nil {
	 	return err
	 }

	 user := rsp.(*entity.Identity)
	 user.Password = newPass
	 _, err = s.repo.Update(user).Unwrap()

	 return err

}



func NewUserService(config *utils.Configuration) UserService {
	repo := db.NewIdentityStore(config)
	query := db.NewIdentityQuery(config)
	return userService{
		repo:  repo,
		query: query,
	}
}

func (s userService) Create(user *entity.Identity) (*entity.Identity, error) {
	n := s.repo.Save(user)
	r, err := n.Unwrap()
	if err != nil {
		return nil, err
	}
	return r.(*entity.Identity), nil
}

func (s userService) Query(query spec.Expression) ([]entity.Identity, error) {
	r := s.query.All(spec.Specify(query))
	data, err := r.Unwrap()
	if err != nil {
		return nil, err
	}

	return data.([]entity.Identity), nil
}

func (s userService) Delete(id int64) error {
	_, e := s.repo.Delete(id).Unwrap()
	if e != nil {
		return e
	}
	return nil
}

func (s userService) Update(user *entity.Identity) error {
	_, e := s.repo.Update(user).Unwrap()
	if e != nil {
		return e
	}
	return nil
}

func (s userService) Single(id int64) (*entity.Identity, error) {

	u, e := s.repo.Single(id).Unwrap()
	if e != nil {
		return nil, e
	}
	return u.(*entity.Identity), nil
}
