package repository

import (
	"github.com/mlambda-net/identity/pkg/domain/spec"
	"github.com/mlambda-net/monads/monad"
)

type IdentityQuery interface {
	Single(spec spec.Spec) monad.Mono
	All(spec spec.Spec) monad.Mono
	Close()
}
