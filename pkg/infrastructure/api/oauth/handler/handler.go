package handler

import (
  "github.com/dgrijalva/jwt-go"
  "github.com/go-oauth2/oauth2/v4"
  "github.com/go-oauth2/oauth2/v4/errors"
  "github.com/go-oauth2/oauth2/v4/manage"
  "github.com/go-oauth2/oauth2/v4/server"
  "github.com/go-oauth2/oauth2/v4/store"
  "github.com/go-session/session"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/conf"
  "github.com/mlambda-net/identity/pkg/infrastructure/db"
  "github.com/mlambda-net/net/pkg/net"
  "log"
  "net/http"
)

type handler struct {
  store *store.ClientStore
  srv   *server.Server
  auth  net.Request
  conf  *conf.Configuration
  cache repository.IdentityCache
}

func NewHandler(auth net.Request, conf *conf.Configuration) *handler {

  appConfig := &utils.AppConfig{}
  appConfig.Cache = conf.Cache
  cache := db.NewClientCache(appConfig)

  manager := manage.NewDefaultManager()
  manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

  manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)
  manager.SetPasswordTokenCfg(manage.DefaultPasswordTokenCfg)

  manager.MustTokenStorage(store.NewMemoryTokenStore())
  manager.MapAccessGenerate(NewClaimAccess("token", conf, cache, jwt.SigningMethodHS512))

  clientStore := GetStore()
  manager.MapClientStorage(clientStore)

  srv := server.NewDefaultServer(manager)
  srv.SetPasswordAuthorizationHandler(Authenticate(cache, auth))
  srv.SetUserAuthorizationHandler(AuthorizeHandler)
  srv.SetClientInfoHandler(server.ClientFormHandler)
  srv.SetAllowedGrantType(oauth2.AuthorizationCode, oauth2.PasswordCredentials, oauth2.ClientCredentials, oauth2.Refreshing)
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
    auth:  auth,
    conf:  conf,
    cache: cache,
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


func Authenticate(cache repository.IdentityCache, auth net.Request) func(username string, password string) (userID string, err error) {

  return func(username string, password string) (userID string, err error) {
    resp, err := auth.Request(&message.Authenticate{
      Email:    username,
      Password: password,
    }).Unwrap()

    if err != nil {
      return "", err
    }

    user := resp.(*message.User)

    cache.Set(user.Id, user)

    return user.Id, nil
  }
}

func setupHeaders(w http.ResponseWriter, r *http.Request) {
  (w).Header().Set("Access-Control-Allow-Origin", "*")
  (w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  (w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  if r.Method == "OPTIONS" {
    w.WriteHeader(http.StatusOK)
    return
  }
}
