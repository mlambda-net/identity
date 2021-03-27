package user

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/application/message"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/net/pkg/ex"
)

func (u *userActor) deleteUser(msg *message.Delete) core.Resolver {
  resolve := core.NewResolve()
  id, e := uuid.Parse( msg.Id )
  if e != nil{
    return resolve.Error(ex.Error(e))
  }
  return resolve.Mono( u.service.Delete(id)).Then(func(_ types.Any) types.Any {
    return &core.Done{}
  })
}
