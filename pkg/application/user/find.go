package user

import (
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
)

func (u *userActor) Find(msg *message.Find) core.Resolver {
  resolve := core.NewResolve()
  return resolve.Mono(u.query.ByEmail(msg.Email)).Then(func(any types.Any) types.Any {
    usr := any.(*entity.Identity)
    return &message.Result{
      Id:       usr.ID.String(),
      Name:     usr.Name,
      LastName: usr.LastName,
      Email:    usr.Email,
    }
  })
}

