package handler

import (
  "context"
  "encoding/base64"
  "github.com/dgrijalva/jwt-go"
  "github.com/go-oauth2/oauth2/v4"
  "github.com/go-oauth2/oauth2/v4/generates"
  "github.com/google/uuid"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/conf"
  "strings"
)

type claimsAccess struct {
  SignedKeyID  string
  SignedKey    []byte
  SignedMethod jwt.SigningMethod
  conf         *conf.Configuration
}


func NewClaimAccess(kid string, conf *conf.Configuration, method jwt.SigningMethod) oauth2.AccessGenerate {
  return &claimsAccess{
    SignedKeyID:  kid,
    conf: conf,
    SignedMethod: method,
  }
}

func (c *claimsAccess) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (access, refresh string, err error)  {
  claims := &generates.JWTAccessClaims{
    StandardClaims: jwt.StandardClaims{
      Audience:  data.Client.GetID(),
      Subject:   data.UserID,
      ExpiresAt: data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix(),
      Issuer: c.conf.OAuth.Host,
    },
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
