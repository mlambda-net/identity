package user

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/net/pkg/ex"
)

func (u *userActor) Create( msg *message.Create) core.Resolver {
  resolve := core.NewResolve()

  roles, e := u.GetRoles(msg.Roles)
  if e != nil {
    resolve.Error(ex.Error(e))
  }
  user := &entity.Identity{
    ID:       uuid.New(),
    Name:     msg.Name,
    LastName: msg.LastName,
    Email:    msg.Email,
    Password: msg.Password,
    Roles:    roles,
  }

  return resolve.Mono(u.service.Create(user)).Then(func(any types.Any) types.Any {
    return &message.Response{Id: user.ID.String()}
  })

}
