package auth

import (
  "errors"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/net/pkg/ex"
)

func (a *authActor) authenticate(msg *message.Authenticate) core.Resolver {
  resolve := core.NewResolve()
  return resolve.Mono(a.service.Authenticate(msg.Email, msg.Password).Bind(func(any types.Any) monad.Mono {
    if !any.(bool) {
      monad.ToMono(ex.Error(errors.New("the user or password are incorrect")))
    }
    return a.query.ByEmail(msg.Email)
  })).Then(func(any types.Any) types.Any {
    user := any.(*entity.Identity)

    var roles []*message.Role
    for _, role := range user.Roles{
      roles = append(roles, &message.Role{
        App:  role.App.Name,
        Name: role.Name,
      })
    }

    return &message.User{
      Id:       user.ID.String(),
      Name:     user.Name,
      LastName: user.LastName,
      Email:    user.Email,
      Roles: roles,
    }
  })
}

