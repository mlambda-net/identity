package app

import (
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/core"
)

func (a *appActor) Find(msg *message.AppFind) core.Resolver {
  return core.NewResolve().Mono(a.query.Find(msg.Id)).Then(func(any types.Any) types.Any {
    switch m := any.(type) {
    case monad.Just:
      app := m.Value().(*entity.App)
      return &message.AppResult{
        Id:          app.ID.String(),
        Name:        app.Name,
        Description: app.Description,
      }
    default:
      return &message.AppNotFound{}
    }
  })
}
