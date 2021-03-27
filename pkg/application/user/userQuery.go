package user

import (
  "errors"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/specs"
  "github.com/mlambda-net/net/pkg/core"
)

func (u *userActor) userQuery(msg *message.Filter) core.Resolver {
  switch msg.By {
  case message.EMAIL:
    return u.filter(specs.ByEmail(msg.Email))
  case message.ID:
    return u.filter(specs.ById(msg.Id))
  case message.NAME:
    return u.filter(specs.ByName(msg.Name, msg.LastName))
  }
  return core.NewResolve().Error(errors.New("this option is not supported"))
}
