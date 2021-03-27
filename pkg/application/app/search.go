package app

import (
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
)

func (a *appActor) Search(_ *message.AppSearch) core.Resolver {

  resolver := core.NewResolve()

  return resolver.Mono(a.query.Search("")).Then(func(any types.Any) types.Any {
    var results []*message.AppResult

    for _, i := range any.([]*entity.App) {
      results = append(results, &message.AppResult{
        Id:          i.ID.String(),
        Name:        i.Name,
        Description: i.Description,
      })
    }
    return &message.AppResults{
      Results: results,
    }
  })
}

