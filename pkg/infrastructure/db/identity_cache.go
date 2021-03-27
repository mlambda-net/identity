package db

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/AsynkronIT/protoactor-go/log"
  "github.com/go-redis/redis/v8"
  "github.com/mlambda-net/identity/pkg/domain/repository"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/sirupsen/logrus"
)


type identityCache struct {
  db *redis.Client
}

func (s identityCache) Set(key string, data interface{}) {
  ctx := context.Background()
  buffer := serialize(data)
  err := s.db.Set(ctx, key, buffer, 0).Err()
  if err != nil {
    logrus.Error(err)
  }
}

func (s identityCache) Get(key string, value interface{}) {
  ctx := context.Background()
  data, err := s.db.Get(ctx, key).Result()
  if err != nil {
    log.Error(err)
    return
  }
  _ = json.Unmarshal([]byte(data), value)
}

func NewCache(config *utils.Configuration) repository.IdentityCache  {
  rdb := initializeCache(config)

  return identityCache{db: rdb}
}



func NewClientCache(config *utils.AppConfig) repository.IdentityCache {
  rdb := redis.NewClient(&redis.Options{
    Addr:     fmt.Sprintf("%s:%d", config.Cache.Server, config.Cache.Port),
    Password: config.Cache.Password,
    DB:       config.Cache.DB,
  })

  return identityCache{db: rdb}
}
