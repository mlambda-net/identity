package app

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/net/pkg/ex"
)

func (a *appActor) Edit( msg *message.AppEdit) core.Resolver {
  resolve := core.NewResolve()

  id, err := uuid.Parse(msg.Id)
  if err != nil {
    return resolve.Error(ex.Error(err))
  }
  return resolve.Mono(a.service.Update(&entity.App{
    ID:          id,
    Description: msg.Description,
  })).Then(func(any types.Any) types.Any {
    return &message.AppId{Id: any.(*entity.App).ID.String()}
  })
}
