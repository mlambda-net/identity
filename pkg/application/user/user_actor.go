package user

import (
  "github.com/AsynkronIT/protoactor-go/actor"
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/identity/pkg/domain/services"
  "github.com/mlambda-net/identity/pkg/domain/utils"
)

type userActor struct {
  service services.UserService
  query   services.QueryService
}

func NewUserProps(config *utils.Configuration) *actor.Props {
	service := services.NewUserService(config)
	query := services.NewQueryService(config)
	return actor.PropsFromProducer(func() actor.Actor { return &userActor{service: service, query: query} })
}

func (u *userActor) Receive(ctx actor.Context) {
  switch msg := ctx.Message().(type) {
  case *message.Create:
    ctx.Respond(u.Create( msg).Response())
  case *message.GetUser:
    ctx.Respond(u.Get(msg).Response())
  case *message.Filter:
    ctx.Respond(u.userQuery(msg).Response())
  case *message.Update:
    ctx.Respond(u.update(msg).Response())
  case *message.Delete:
    ctx.Respond(u.deleteUser(msg).Response())
  case *message.ChangePassword:
    ctx.Respond(u.changePassword(msg).Response())
  case *message.Find:
    ctx.Respond(u.Find(msg).Response())
  case *message.All:
    ctx.Respond(u.All(msg).Response())
  }
}



func (u *userActor) GetRoles(ids []string) ([]*entity.Role, error) {
  roles := make([]*entity.Role,0)
  for _, r := range ids {
    id, e := uuid.Parse(r)
    if e != nil {
      return nil, e
    }
    roles = append(roles, &entity.Role{
      ID: id,
    })
  }
  return roles, nil
}
