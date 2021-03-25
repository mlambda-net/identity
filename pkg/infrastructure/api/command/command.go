package command

import (
  "fmt"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/conf"
  "github.com/mlambda-net/net/pkg/net"
  "golang.org/x/oauth2"
)

type Command interface {
  Register(r net.Route)
}

type control struct {
  user  net.Request
  auth  net.Request
  oauth oauth2.Config
}

func NewCommand(user, auth net.Request, conf *conf.Configuration) Command {

  config := oauth2.Config{
    ClientID:     conf.OAuth.ClientID,
    ClientSecret: conf.OAuth.Secret,
    Endpoint: oauth2.Endpoint{
      AuthURL:  fmt.Sprintf("%s/authorize", conf.OAuth.Host),
      TokenURL: fmt.Sprintf("%s/token", conf.OAuth.Host),
    },
    RedirectURL: fmt.Sprintf("%s/identity/oauth2", conf.App.Url),
    Scopes:      []string{ "admin", "reader", "all"},
  }

  return &control{
    user: user,
    auth: auth,
    oauth: config,
  }
}



func (c *control) Register(r net.Route) {
  r.AddRoute("root", "/identity/", false, "GET", c.root)
  r.AddRoute("token","/identity/oauth2", false, "GET", c.OAuth2)
  r.AddRoute("token","/identity/login", false, "POST", c.login)
  r.AddRoute("register", "/identity/user", false, "POST", c.createUser)
  r.AddRoute("user_update", "/identity/user", true, "PUT", c.updateUser)
  r.AddRoute("user_delete", "/identity/user", true, "DELETE", c.deleteUser)
  r.AddRoute("user_find", "/identity/user/{email}", true, "GET", c.getUser)
  r.AddRoute("user_change_password", "/identity/user/change_password", true, "POST", c.changePassword)
}
