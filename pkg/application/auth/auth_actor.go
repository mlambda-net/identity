package auth

import (
  "github.com/AsynkronIT/protoactor-go/actor"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/services"
  "github.com/mlambda-net/identity/pkg/domain/utils"
)

type authActor struct {
  service services.UserService
  query   services.QueryService
}

func NewAuthProps(config *utils.Configuration) *actor.Props {
	service := services.NewUserService(config)
	query := services.NewQueryService(config)
	return actor.PropsFromProducer(func() actor.Actor { return &authActor{service: service, query: query} })
}

func (a *authActor) Receive(ctx actor.Context) {
  switch msg := ctx.Message().(type) {
  case *message.Register:
    ctx.Respond(a.register(msg).Response())
  case *message.Authenticate:
    ctx.Respond(a.authenticate(msg).Response())
  }
}
