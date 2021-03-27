package roles

import (
  "github.com/AsynkronIT/protoactor-go/actor"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/services"
  "github.com/mlambda-net/identity/pkg/domain/utils"
)

type rolesActor struct {
  service services.RolesService
  query   services.RolesQuery
}

func (r rolesActor) Receive(ctx actor.Context) {
  switch msg := ctx.Message().(type) {
  case *message.RoleAdd:
    ctx.Respond( r.add(msg).Response())
  case *message.RoleEdit:
   ctx.Respond( r.update(msg).Response())
  case *message.RoleDelete:
    ctx.Respond( r.delete( msg).Response())
  case *message.RoleGet:
    ctx.Respond( r.get(msg).Response())
  case *message.RoleSearch:
    ctx.Respond(r.search(msg).Response())
  }
}

func NewRolesProps(config *utils.Configuration) *actor.Props {
  service := services.NewRolesService(config)
  query := services.NewRolesQuery(config)
  return actor.PropsFromProducer(func() actor.Actor { return &rolesActor{service: service, query: query} })
}
