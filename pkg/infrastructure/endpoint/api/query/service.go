package query

import (
  "github.com/graphql-go/graphql"
  "github.com/mlambda-net/identity/pkg/application/message"
)

func (c *control) getUsers(token string) func (params graphql.ResolveParams) (interface{}, error) {
  return func(params graphql.ResolveParams) (interface{}, error) {
    email, ok := params.Args["email"].(string)
    if ok {
      return c.fetchUsers(token, email)
    }
    return message.Result{}, nil
  }

}

func (c *control) fetchUsers(token, email string) ([]*message.Result, error) {
  result, err := c.user.Token(token).Request(&message.Filter{
    Email: email,
    By:    message.EMAIL,
  }).Unwrap()

  if err != nil {
    return nil, err
  }

  rs := result.(*message.Results)
  return rs.Results, nil
}
