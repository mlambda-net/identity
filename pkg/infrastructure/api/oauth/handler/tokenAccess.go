package handler

import (
  "context"
  "encoding/base64"
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "github.com/go-oauth2/oauth2/v4"
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/application/message"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/conf"
  "strings"
)

type claimsAccess struct {
  SignedKeyID  string
  SignedKey    []byte
  SignedMethod jwt.SigningMethod
  conf         *conf.Configuration
  cache        repository.IdentityCache
}


func NewClaimAccess(kid string, conf *conf.Configuration, cache repository.IdentityCache, method jwt.SigningMethod) oauth2.AccessGenerate {
  return &claimsAccess{
    SignedKeyID:  kid,
    conf:         conf,
    SignedMethod: method,
    cache:        cache,
  }
}

func (c *claimsAccess) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (access, refresh string, err error)  {
  var user *message.User
  c.cache.Get(data.UserID, &user)

  roles := make( []role,0)
  for _, r := range user.Roles{
    roles = append(roles, role{
      App:  r.App,
      Name: r.Name,
    })
  }

  claims := &claims{
      Audience:  data.Client.GetID(),
      Subject:   data.UserID,
      ExpiresAt: data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix(),
      Issuer: c.conf.OAuth.Host,
      Email: user.Email,
      Name:  fmt.Sprintf("%s %s", user.Name, user.LastName),
      Roles: roles,
  }

  token := jwt.NewWithClaims(c.SignedMethod, claims)
  if c.SignedKeyID != "" {
    token.Header["kid"] = c.SignedKeyID
  }

  var key = []byte(c.conf.App.Secret)
  access, err = token.SignedString(key)
  if err != nil {
    return "", "", err
  }
  refresh = ""

  if isGenRefresh {
    t := uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(access)).String()
    refresh = base64.URLEncoding.EncodeToString([]byte(t))
    refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
  }

  return access, refresh, nil

}
