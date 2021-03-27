package db

import (
  "encoding/json"
  "fmt"
  "github.com/go-pg/pg/v10"
  "github.com/go-redis/redis/v8"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/sirupsen/logrus"
)

func  serialize(obj interface{}) []byte  {
  b, err := json.Marshal(obj)
  if err != nil {
    logrus.Error(err)
    return nil
  }
  return b
}


func initializeCache(config *utils.Configuration) *redis.Client {
  rdb := redis.NewClient(&redis.Options{
    Addr:     fmt.Sprintf("%s:%d", config.Cache.Server, config.Cache.Port),
    Password: config.Cache.Password,
    DB:       config.Cache.DB,
  })
  return rdb
}

func initializeDB(config *utils.Configuration) *pg.DB {
  db := pg.Connect(&pg.Options{
    User:     config.Db.User,
    Password: config.Db.Password,
    Addr:     fmt.Sprintf("%s:%s", config.Db.Host, config.Db.Port),
    Database: config.Db.Schema,
  })
  return db
}

