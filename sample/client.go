package main

import (
  "context"
  "golang.org/x/oauth2"
)

func main() {

  config := oauth2.Config{
    ClientID:     "222222",
    ClientSecret: "22222222",
    Scopes:       []string{"all"},
    RedirectURL:  "http://localhost:9094/oauth2",
    Endpoint: oauth2.Endpoint{
      AuthURL:  "http://localhost:8000/identity/authorize",
      TokenURL: "http://localhost:8000/identity/token",
    },
  }

  token, _ :=  config.Exchange(context.Background(),"code")

  println(token)

}


