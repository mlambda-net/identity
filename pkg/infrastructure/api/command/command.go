package command

import (
  "fmt"
  "github.com/gorilla/mux"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/conf"
  "github.com/mlambda-net/net/pkg/net"
  "github.com/mlambda-net/net/pkg/security"
  "golang.org/x/oauth2"
  "net/http"
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
  r.AddRoute("root", "/identity", false, "GET", c.root)
  r.AddRoute("token", "/identity/oauth2", false, "GET", c.OAuth2)
  r.AddRoute("login", "/identity/login", false, "POST", c.login)

  r.AddRoute("user_create", "/identity/user", true, "POST", c.createUser)
  r.AddRoute("user_update", "/identity/user", true, "PUT", c.updateUser)
  r.AddRoute("user_delete", "/identity/user", true, "DELETE", c.deleteUser)
  r.AddRoute("user_get", "/identity/user/{id}", true, "GET", c.getUser)
  r.AddRoute("user_change_password", "/identity/user/change_password", true, "POST", c.changePassword)
  //r.AddRoute("users_get", "/identity/user/{filter}", true, "GET", c.getUsers)



  r.AddRoute("app_add", "/identity/app", true, "POST", c.appAdd)
  r.AddRoute("app_edit", "/identity/app", true, "PUT", c.appEdit)
  r.AddRoute("app_get", "/identity/app/{id}", true, "GET", c.appGet)
  r.AddRoute("app_all", "/identity/app", true, "GET", c.appGetAll)

  r.AddRoute("role_add", "/identity/role", true, "POST", c.roleAdd)
  r.AddRoute("role_edit", "/identity/role", true, "PUT", c.roleEdit)
  r.AddRoute("role_delete", "/identity/role", true, "DELETE", c.roleDelete)
  r.AddRoute("role_get", "/identity/role/{id}", true, "GET", c.roleGet)
  r.AddRoute("role_all", "/identity/role", true, "GET", c.roleAll)

  r.AddRoute("register", "/identity/register", false, "POST", c.register)
  r.AddRoute("profile", "/identity/profile", true, "GET", c.profile)



  r.Add(func(router *mux.Router) {
    router.Handle("/identity/user",security.Authenticate(http.HandlerFunc(c.getUsers))).Methods("GET").Queries("filter", "{filter}")
  })
}
