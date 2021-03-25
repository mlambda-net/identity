package services

import (
  "errors"
  "fmt"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/mlambda-net/identity/pkg/infrastructure/db"
  "github.com/mlambda-net/net/pkg/security"
  "os"
)

type UserService interface {
	Create(user *entity.Identity) (*entity.Identity, error)
	Delete(id int64) error
	Update(user *entity.Identity) error
	ChangePassword(email string, password string) error
	Authenticate(login string, password string) (string, error)
}

type userService struct {
	repo  repository.IdentityStore
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





func (s userService) Create(user *entity.Identity) (*entity.Identity, error) {
	n := s.repo.Save(user)
	r, err := n.Unwrap()
	if err != nil {
		return nil, err
	}
	return r.(*entity.Identity), nil
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



func NewUserService(config *utils.Configuration) UserService {
  repo := db.NewIdentityStore(config)
  return userService{
    repo:  repo,
  }
}
