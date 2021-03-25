package server

import (
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/sirupsen/logrus"
  "os"
  "strconv"
)

func (s *server) LoadConfig() {
  s.config = &utils.Configuration{}
  s.config.Env = os.Getenv("ENV")

  s.config.App.Name = os.Getenv("APP_NAME")
  s.config.App.Port = os.Getenv("APP_PORT")
  s.config.App.Version = os.Getenv("1.0.0")

  s.config.Metric.Namespace = os.Getenv("METRIC_NAMESPACE")
  s.config.Metric.Port = os.Getenv("METRIC_PORT")

  s.config.Email.Server = os.Getenv("EMAIL_SERVER")
  s.config.Email.Port = toint(os.Getenv("EMAIL_PORT"))

  s.config.Db.Host = os.Getenv("DB_HOST")
  s.config.Db.User = os.Getenv("DB_USER")
  s.config.Db.Password = os.Getenv("DB_PASSWORD")
  s.config.Db.Port = os.Getenv("DB_PORT")
  s.config.Db.Schema = os.Getenv("DB_DATA")

  s.config.Cache.Server = os.Getenv("CACHE_SERVER")
  s.config.Cache.Password = os.Getenv("CACHE_PASSWORD")
  s.config.Cache.DB = toint(os.Getenv("CACHE_DB"))
  s.config.Cache.Port = toint(os.Getenv("CACHE_PORT"))

  s.config.Index.Server = os.Getenv("INDEX_SERVER")
  s.config.Index.Authenticate = tobool(os.Getenv("INDEX_AUTHENTICATE"))
  s.config.Index.User = os.Getenv("INDEX_USER")
  s.config.Index.Password = os.Getenv("INDEX_PASSWORD")

}

func toint(value string) int {
  rtn, err := strconv.Atoi( value)
  if err != nil {
    logrus.Fatal(err)
    return 0
  }
  return rtn
}

func tobool(value string) bool  {
  rtn, err := strconv.ParseBool("true")
  if err != nil {
    logrus.Fatal(err)
    return false
  }
  return rtn
}
