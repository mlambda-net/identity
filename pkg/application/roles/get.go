package roles

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/net/pkg/ex"
)

func (r rolesActor) get(msg *message.RoleGet) core.Resolver {
  resolve := core.NewResolve()

  id, err := uuid.Parse(msg.Id)
  if err != nil {
    return resolve.Error(ex.Error(err))
  }

  return resolve.Mono(r.query.Find(id)).Then(func(any types.Any) types.Any {
    switch m := any.(type) {
    case monad.Just:
      rol := m.Value().(*entity.Role)
      return &message.RoleResult{
        Id:          rol.ID.String(),
        App:         rol.App.ID.String(),
        AppName:     rol.App.Name,
        Name:        rol.Name,
        Description: rol.Description,
      }
    default:
      return &message.RoleNotFound{}
    }
  })
}

