package user

import (
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/entity"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
)

func (u *userActor) All(msg *message.All) core.Resolver {
  resolve := core.NewResolve()
  return resolve.Mono(u.query.All(msg.Filter)).Then(func(any types.Any) types.Any {
    items := any.([]*entity.Identity)
    var users []*message.Result

    for _, item := range items {
      users = append(users, &message.Result{
        Id:       item.ID.String(),
        Name:     item.Name,
        LastName: item.LastName,
        Email:    item.Email,
      })
    }
    return &message.Results{Results: users}
  })
}
