package utils

type Configuration struct {
  Env string `default:"dev"`

  App struct {
    Name    string
    Version string `default:"1.0.0"`
    Port    string `default:"9001"`
  }

  Metric struct {
    Port      string `default:"9002"`
    Namespace string `default:"ns"`
  }

  Email struct{
    Server string
    Port  int
  }

  Db struct {
    Host     string `default:"localhost"`
    Port     string `default:"5432"`
    User     string `default:"postgres"`
    Password string `default:"123"`
    Schema   string `default:"postgres"`
  }

  Cache struct {
    Server   string `default:"localhost"`
    Port     int    `default:"6379"`
    Password string `default:""`
    DB       int    `default:"0"`
  }

  Index struct{
    Server string
    Authenticate bool
    User string
    Password string
  }

}
