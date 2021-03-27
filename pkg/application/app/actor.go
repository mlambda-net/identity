package app

import (
  "github.com/AsynkronIT/protoactor-go/actor"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/services"
  "github.com/mlambda-net/identity/pkg/domain/utils"
)

type appActor struct {
  service services.AppService
  query   services.AppQuery
}

func (a *appActor) Receive(ctx actor.Context) {
  switch msg := ctx.Message().(type) {
  case *message.AppAdd:
    ctx.Respond(a.Add(msg).Response())

  case *message.AppEdit:
    ctx.Respond(a.Edit(msg).Response())

  case *message.AppSearch:
    ctx.Respond(a.Search(msg).Response())

  case *message.AppFind:
    ctx.Respond(a.Find(msg).Response())
  }
}

func NewAppActor(config *utils.Configuration) *actor.Props  {
  service := services.NewAppService(config)
  query := services.NewAppQuery(config)
  return actor.PropsFromProducer(func() actor.Actor {
    return &appActor{service: service, query: query}
  })
}
