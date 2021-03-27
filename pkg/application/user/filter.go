package user

import (
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/net/pkg/spec"
)

func (u *userActor) filter(spec spec.Expression) core.Resolver {
  resolve := core.NewResolve()

  return resolve.Mono(u.query.Query(spec)).Then(func(any types.Any) types.Any {
    results := make([]*message.Result, 0)
    for _, user := range any.([]*entity.Identity) {
      r := &message.Result{
        Id:       user.ID.String(),
        Name:     user.Name,
        LastName: user.LastName,
        Email:    user.Email,
      }
      results = append(results, r)
    }
    return &message.Results{Results: results}
  })
}
