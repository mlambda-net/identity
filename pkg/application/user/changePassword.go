package user

import (
  "errors"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/net/pkg/ex"
)

func (u *userActor) changePassword(msg *message.ChangePassword) core.Resolver {
  resolve := core.NewResolve()
  return resolve.Mono(u.query.ByEmail(msg.Email).Bind(func(any types.Any) monad.Mono {
    usr := any.(*entity.Identity)
    old, e := utils.Decrypt(usr.Password)
    if e != nil {
      return monad.ToMono(ex.Error(e))
    }
    if msg.Password != old {
      return monad.ToMono(ex.Error(errors.New("password mismatch")))
    }
    return u.service.ChangePassword(msg.Email, msg.Password)
  })).Then(func(_ types.Any) types.Any {
    return core.Done{}
  })
}
