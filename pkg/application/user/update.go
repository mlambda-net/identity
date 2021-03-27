package user

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/net/pkg/ex"
)

func (u *userActor) update(msg *message.Update) core.Resolver {

  resolve := core.NewResolve()

  id, e := uuid.Parse(msg.Id)
  if e != nil {
    return resolve.Error(ex.Error(e))
  }

  roles, e := u.GetRoles(msg.Roles)
  if e != nil {
    resolve.Error(ex.Error(e))
  }

  return resolve.Mono(u.query.Single(id).Bind(func(any types.Any) monad.Mono {
    usr := any.(*entity.Identity)
    usr.Name = msg.Name
    usr.LastName = msg.LastName
    usr.Roles = roles
    return u.service.Update(usr)
  })).Then(func(any types.Any) types.Any {
    return &message.Response{Id: any.(*entity.Identity).ID.String()}
  })
}
