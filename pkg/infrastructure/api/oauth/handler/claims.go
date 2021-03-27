package handler

import "github.com/dgrijalva/jwt-go"

type role struct {
  App string `json:"app"`
  Name string `json:"name"`
}

type claims struct {
  Audience  string   `json:"aud,omitempty"`
  ExpiresAt int64    `json:"exp,omitempty"`
  Id        string   `json:"jti,omitempty"`
  IssuedAt  int64    `json:"iat,omitempty"`
  Issuer    string   `json:"iss,omitempty"`
  NotBefore int64    `json:"nbf,omitempty"`
  Subject   string   `json:"sub,omitempty"`
  Email     string   `json:"email,omitempty"`
  Name      string   `json:"name,omitempty"`
  Roles     []role `json:"roles"`
}

func (c claims) Valid() error {
  standard := &jwt.StandardClaims{
    Audience:  c.Audience,
    ExpiresAt: c.ExpiresAt,
    Id:        c.Id,
    IssuedAt:  c.IssuedAt,
    Issuer:    c.Issuer,
    NotBefore: c.NotBefore,
    Subject:   c.Subject,
  }
  return standard.Valid()
}

