package handler

import (
  "github.com/go-oauth2/oauth2/v4/models"
  "github.com/go-oauth2/oauth2/v4/store"
)

func GetStore() *store.ClientStore {
  clientStore := store.NewClientStore()
  clientStore.Set("abc", &models.Client{
    ID:     "abc",
    Secret: "123",
    Domain: "http://localhost:3000",
  })

  clientStore.Set("identity", &models.Client{
    ID:     "identity",
    Secret: "123",
    Domain: "http://localhost",
  })

  clientStore.Set("swagger", &models.Client{
    ID:     "swagger",
    Secret: "123",
    Domain: "http://localhost",
  })
  return clientStore
}
