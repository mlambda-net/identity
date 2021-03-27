package oauth

import (
  "fmt"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/conf"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/oauth/handler"
  "github.com/mlambda-net/net/pkg/net"
  log "github.com/sirupsen/logrus"
  "net/http"
)

type OAuth interface {
  Start()
}

type oAuth struct {
  conf  *conf.Configuration
  auth  net.Request
}

func NewOAuthServer(conf *conf.Configuration, auth net.Request) OAuth {
  return &oAuth{
    auth: auth,
    conf : conf,
  }
}


func (o *oAuth) Start() {

  go func() {
    h := handler.NewHandler(o.auth, o.conf)
    http.HandleFunc("/login", h.Login)
    http.HandleFunc("/auth", h.Authenticate)
    http.HandleFunc("/authorize", h.Authorize)
    http.HandleFunc("/token", h.Token)
    http.HandleFunc("/fail", h.Fail)
    http.HandleFunc("/.well-known/openid-configuration", h.WellKnow)
    log.Printf("OAuth is running at %s port.", o.conf.App.OAuthPort)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", o.conf.App.OAuthPort), nil))
  }()
}
