package handler

import (
  "github.com/dgrijalva/jwt-go"
  "github.com/go-oauth2/oauth2/v4/errors"
  "github.com/go-oauth2/oauth2/v4/manage"
  "github.com/go-oauth2/oauth2/v4/models"
  "github.com/go-oauth2/oauth2/v4/server"
  "github.com/go-oauth2/oauth2/v4/store"
  "github.com/go-session/session"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/conf"
  "github.com/mlambda-net/net/pkg/net"
  "log"
  "net/http"
)

type handler struct {
  store *store.ClientStore
  srv   *server.Server
  auth  net.Request
  conf  *conf.Configuration
}

func NewHandler(auth net.Request, conf *conf.Configuration) *handler {

  manager := manage.NewDefaultManager()
  manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
  manager.MustTokenStorage(store.NewMemoryTokenStore())

  manager.MapAccessGenerate(NewClaimAccess("", conf, jwt.SigningMethodHS512))

  clientStore := store.NewClientStore()
  manager.MapClientStorage(clientStore)

  clientStore.Set("abc", &models.Client{
    ID:     "identity",
    Secret: "123",
    Domain: "http://localhost:8000",
  })

  clientStore.Set("swagger", &models.Client{
    ID:     "swagger",
    Secret: "123",
    Domain: "http://localhost:8002",
  })


  srv := server.NewDefaultServer(manager)
  srv.SetPasswordAuthorizationHandler(Authenticate(auth))
  srv.SetUserAuthorizationHandler(AuthorizeHandler)
  srv.SetAllowGetAccessRequest(true)


  manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

  srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
    log.Println("Internal Error:", err.Error())
    return
  })

  srv.SetResponseErrorHandler(func(re *errors.Response) {
    log.Println("Response Error:", re.Error.Error())
  })

  return &handler{
    store: clientStore,
    srv:   srv,
    auth: auth,
    conf : conf,
  }

}

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
  store, err := session.Start(r.Context(), w, r)
  if err != nil {
    return
  }

  uid, ok := store.Get("LoggedInUserID")
  if !ok {
    if r.Form == nil {
      r.ParseForm()
    }

    store.Set("ReturnUri", r.Form)
    store.Save()

    w.Header().Set("Location", "/login")
    w.WriteHeader(http.StatusFound)
    return
  }

  userID = uid.(string)
  store.Delete("LoggedInUserID")
  store.Save()
  return
}


func Authenticate( auth net.Request) func (username string, password string) (userID string, err error) {

  return func(username string, password string) (userID string, err error) {
    _, err = auth.Request(&message.Authenticate{
      Email:    username,
      Password: password,
    }).Unwrap()

    if err != nil {
      return "", err
    }

    return username, nil
  }

}
