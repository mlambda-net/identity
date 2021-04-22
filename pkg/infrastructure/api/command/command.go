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
  oauth oauth2.Config
  user  net.Request
  auth  net.Request
  app   net.Request
  role  net.Request
}

func NewCommand(user, auth, app, role net.Request, conf *conf.Configuration) Command {
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
    app: app,
    role: role,
    oauth: config,
  }
}

func (c *control) Register(r net.Route) {
  r.AddRoute("root", "/identity", nil,false, "GET", c.root)
  r.AddRoute("token", "/identity/oauth2",nil, false, "GET", c.OAuth2)
  r.AddRoute("login", "/identity/login",nil, false, "POST", c.login)

  r.AddRoute("user_create", "/identity/user", nil,true, "POST", c.createUser)
  r.AddRoute("user_update", "/identity/user", nil,true, "PUT", c.updateUser)
  r.AddRoute("user_delete", "/identity/user", nil,true, "DELETE", c.deleteUser)
  r.AddRoute("user_get", "/identity/user/{id}",nil, true, "GET", c.getUser)
  r.AddRoute("user_change_password", "/identity/user/change_password",nil, true, "POST", c.changePassword)
  r.AddRoute("user_all", "/identity/user",[]string {"filter"} , true, "GET", c.getUsers)

  r.AddRoute("app_add", "/identity/app",nil, true, "POST", c.appAdd)
  r.AddRoute("app_edit", "/identity/app",nil, true, "PUT", c.appEdit)
  r.AddRoute("app_get", "/identity/app/{id}", nil,true, "GET", c.appGet)
  r.AddRoute("app_all", "/identity/app",nil, true, "GET", c.appGetAll)

  r.AddRoute("role_add", "/identity/role",nil, true, "POST", c.roleAdd)
  r.AddRoute("role_edit", "/identity/role",nil, true, "PUT", c.roleEdit)
  r.AddRoute("role_delete", "/identity/role",nil, true, "DELETE", c.roleDelete)
  r.AddRoute("role_get", "/identity/role/{id}",nil, true, "GET", c.roleGet)
  r.AddRoute("role_all", "/identity/role",nil, true, "GET", c.roleAll)

  r.AddRoute("register", "/identity/register",nil, false, "POST", c.register)
  r.AddRoute("profile", "/identity/profile",nil, true, "GET", c.profile)

}
