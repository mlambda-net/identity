package db

import (
  "context"
  "github.com/mlambda-net/identity/pkg/domain/utils"
)

func AliveDB(config *utils.Configuration) error {
  db:= initializeDB(config)
  return db.Ping(context.Background())
}

func AliveCache(config *utils.Configuration) error {
  cache := initializeCache(config)
  _, err := cache.Ping(context.Background()).Result()
  return err
}

