package query

import (
  "github.com/mlambda-net/net/pkg/net"
)

type Query interface {
  Register(r net.Route)
}

type control struct {
  user   net.Request
}

func (c *control) Register(r net.Route) {
  r.AddRoute("graphql", "/api/graphql", nil,true, "POST", c.handler)
}

func NewQuery( user net.Request) Query  {
  return &control{user: user}
}
