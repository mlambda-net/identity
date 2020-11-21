package command

import (
  "github.com/mlambda-net/net/pkg/net"
)

type Command interface {
  Register(r net.Route)
}

type control struct {
	user   net.Request
	auth   net.Request
}

func NewCommand(user, auth net.Request) Command {
  return &control{
    user: user,
    auth: auth,
  }
}

func (c *control) Register(r net.Route) {
  r.AddRoute("login", "/user/auth", false, "POST", c.loginUser)
  r.AddRoute("register", "/user/user", false, "POST", c.createUser)
  r.AddRoute("user_update", "/user/user", true, "PUT", c.updateUser)
  r.AddRoute("user_delete", "/user/user", true, "DELETE", c.deleteUser)
  r.AddRoute("user_find", "/user/user/{email}", true, "GET", c.getUser)
  r.AddRoute("user_change_password", "/user/user/change_password", true, "POST", c.changePassword)
}
