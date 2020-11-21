package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/mlambda-net/identity/pkg/domain/entity"
	"github.com/mlambda-net/identity/pkg/domain/repository"
	"github.com/mlambda-net/identity/pkg/domain/spec"
	"github.com/mlambda-net/identity/pkg/domain/utils"
	"github.com/mlambda-net/monads/monad"
)

type identityQuery struct {
	db *pg.DB
}

func (i identityQuery) All(spec spec.Spec) monad.Mono {
	var items []entity.Identity

	_, err := i.db.Query(&items, fmt.Sprintf("SELECT * FROM identities where %s", spec.Query()))

	if err != nil {
		monad.ToMono(err)
	}

	return monad.ToMono(items)
}

func (i identityQuery) Single(spec spec.Spec) monad.Mono {
	var items []entity.Identity
	_, err := i.db.Query(&items, fmt.Sprintf("SELECT * FROM identities where %s", spec.Query()))

	if err != nil {
		monad.ToMono(err)
	}

	if len(items) > 0 {
		return monad.ToMono(items[0])
	}

	return monad.ToMono(nil)
}

func (i identityQuery) Close() {
	_ = i.db.Close()
}

func NewIdentityQuery(config *utils.Configuration) repository.IdentityQuery {
	db := pg.Connect(&pg.Options{
		User:     config.Db.User,
		Password: config.Db.Password,
		Addr:     fmt.Sprintf("%s:%s", config.Db.Host, config.Db.Port),
		Database: config.Db.Schema,
	})

	return &identityQuery{db: db}
}
