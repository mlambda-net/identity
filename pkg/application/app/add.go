package app

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
)

func (a *appActor) Add(msg *message.AppAdd) core.Resolver {
  resolve := core.NewResolve()
  return resolve.Mono(a.service.Add(&entity.App{
    ID:          uuid.New(),
    Name:        msg.Name,
    Description: msg.Description,
  })).Then(func(any types.Any) types.Any {
    return &message.AppId{Id: any.(*entity.App).ID.String()}
  })
}

