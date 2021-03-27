package auth

import (
  "errors"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/net/pkg/ex"
)

func (a *authActor) register(msg *message.Register) core.Resolver {
  resolve := core.NewResolve()
  if msg.Password == "" {
    return resolve.Error(ex.Error(errors.New("the password can not be empty")))
  }
  user := entity.NewIdentityFromRegister(msg.Name, msg.LastName, msg.Email)

  return resolve.Mono( a.service.Create(user)).Then(func(any types.Any) types.Any {
    return &message.Response{Id: any.(*entity.Identity).ID.String()}
  })
}

