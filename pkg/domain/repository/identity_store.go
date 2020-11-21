package repository

import (
	"github.com/mlambda-net/identity/pkg/domain/entity"
	"github.com/mlambda-net/monads/monad"
)

type IdentityStore interface {
	Save(id *entity.Identity) monad.Mono
	Delete(id int64) monad.Mono
	Update(user *entity.Identity) monad.Mono
	Close()
	Single(id int64) monad.Mono
	ByEmail(email string) monad.Mono
}
