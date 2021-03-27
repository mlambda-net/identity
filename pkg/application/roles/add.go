package roles

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
)

func (r rolesActor) add(msg *message.RoleAdd) core.Resolver {
  resolve := core.NewResolve()
  id, e := uuid.Parse(msg.App)
  if e != nil {
    return resolve.Error(e)
  }

  return resolve.Mono(r.service.Add(&entity.Role{
    ID:          uuid.New(),
    Name:        msg.Name,
    Description: msg.Description,
    App:         entity.App{ID: id},
  })).Then(func(any types.Any) types.Any {
    return &message.RoleId{Id: any.(*entity.Role).ID.String()}
  })
}
