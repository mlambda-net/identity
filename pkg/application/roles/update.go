package roles

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/net/pkg/ex"
)

func (r rolesActor) update( msg *message.RoleEdit) core.Resolver {

  resolve := core.NewResolve()

  id, e := uuid.Parse(msg.Id)
  if e != nil {
    return resolve.Error(ex.Error(e))
  }

  m := r.service.Update(&entity.Role{
    ID:          id,
    Description: msg.Description,
  })

  return resolve.Mono(m).Then(func(any types.Any) types.Any {
    return &message.RoleId{Id: id.String()}
  })
}
