package handler

import (
  "github.com/go-oauth2/oauth2/v4/models"
  "github.com/go-oauth2/oauth2/v4/store"
)

func GetStore() *store.ClientStore {
  clientStore := store.NewClientStore()

  clientStore.Set("identity", &models.Client{
    ID:     "identity",
    Secret: "xXBqrnOokTId8IOj",
    Domain: "https://identity.mitienda.co.cr",
  })

  clientStore.Set("identityapi", &models.Client{
    ID:     "identityapi",
    Secret: "wY91HaUBggRGdL70",
    Domain: "https://api.mitienda.co.cr",
  })

  clientStore.Set("abc", &models.Client{
    ID:     "abc",
    Secret: "123",
    Domain: "http://localhost:3000",
  })

  clientStore.Set("localhost", &models.Client{
    ID:     "localhost",
    Secret: "123",
    Domain: "http://localhost:8002",
  })

  clientStore.Set("swagger", &models.Client{
    ID:     "swagger",
    Secret: "123",
    Domain: "http://localhost",
  })
  return clientStore
}
