package user

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
)

func (u *userActor) Get(msg *message.GetUser) core.Resolver {
  resolve := core.NewResolve()

  id, e := uuid.Parse(msg.Id)
  if e != nil {
    return resolve.Error(e)
  }

  return resolve.Mono(u.query.Single(id)).Then(func(any types.Any) types.Any {
    user := any.(*entity.Identity)
    return &message.Result{
      Id:       user.ID.String(),
      Name:     user.Name,
      LastName: user.LastName,
      Email:    user.Email,
      Roles:    user.RoleIds(),
    }
  })
}
