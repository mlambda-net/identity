package roles

import (
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
)

func (r rolesActor) search( msg *message.RoleSearch) core.Resolver{
  resolve := core.NewResolve()
  return resolve.Mono( r.query.Search(msg.Filter)).Then(func(any types.Any) types.Any {
    var results []*message.RoleResult
    for _, rol := range any.([]*entity.Role) {
      results = append(results, &message.RoleResult{
        Id:          rol.ID.String(),
        App:         rol.App.ID.String(),
        AppName:     rol.App.Name,
        Name:        rol.Name,
        Description: rol.Description,
      })
    }
    return &message.RoleResults{Results: results}
  })
}
