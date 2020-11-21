package query

import (
  "errors"
  "fmt"
  "github.com/graphql-go/graphql"
  "github.com/mlambda-net/identity/pkg/application/message"
)

func (c *control) userType() *graphql.Object {

  return graphql.NewObject(graphql.ObjectConfig{
    Name: "user",
    Fields: graphql.Fields{
      "id": &graphql.Field{
        Name: "Id",
        Type: graphql.Int,
      },
      "name": &graphql.Field{
        Name: "Name",
        Type: graphql.String,
      },
      "lastname": &graphql.Field{
        Name: "LastName",
        Type: graphql.String,
      },
      "email": &graphql.Field{
        Name: "Email",
        Type: graphql.String,
      },
      "fullName": &graphql.Field{
        Type: graphql.String,

        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          if user, ok := p.Source.(*message.Result); ok {
            return fmt.Sprintf("%s %s", user.Name, user.LastName), nil
          }
          return nil, errors.New("the type is not correct")
        },
      },
    },
  })
}

func (c *control) userQuery(token string) *graphql.Object {

  return graphql.NewObject(graphql.ObjectConfig{
    Name: "users",
    Fields: graphql.Fields{
      "user": &graphql.Field{
        Type:    graphql.NewList(c.userType()),
        Args:    c.ByEmail(),
        Resolve: c.getUsers(token),
      },
    },
  })

}

func (c *control) ByEmail() graphql.FieldConfigArgument {
  return graphql.FieldConfigArgument{
    "email": &graphql.ArgumentConfig{
      Type: graphql.String,
    },
  }
}

func (c control) schema(token string) graphql.Schema  {
  var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: c.userQuery(token),
  })
  return schema
}

func (c *control) exec(query string, token string)  (*graphql.Result, error) {
  result := graphql.Do(graphql.Params{
    Schema:        c.schema(token),
    RequestString: query,
  })
  if len(result.Errors) > 0 {
    return nil, fmt.Errorf("wrong result, unexpected errors: %v", result.Errors)
  }
  return result, nil
}
