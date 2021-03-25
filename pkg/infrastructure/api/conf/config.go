package conf

import (
  "github.com/sirupsen/logrus"
  "os"
  "strconv"
)

type Configuration struct {
  Env string `default:"dev"`

  App struct {
    Name      string
    Version   string `default:"1.0.0"`
    Port      int    `default:"8080"`
    Host      string
    Base      string
    OAuthPort string
    Secret    string
    Url       string
  }

  OAuth struct {
    Host     string
    Web      string
    ClientID string
    Secret   string
  }

  Docs struct{
    Host    string `default:"localhost"`
    Path    string `default:""`
    Port    int `default:"8082"`
  }

  Metric struct {
    Port      int `default:"8081"`
    Namespace string `default:"ns"`
  }

  Local struct {
    Port string `default:"9001"`
    Host string `default:"localhost"`
  }

  Remote struct {
    Port   string `default:"8000"`
    Server string `default:"localhost"`
  }
}

func LoadConfig() *Configuration {

  config := &Configuration{}
  config.Env = os.Getenv("ENV")

  config.App.Name = os.Getenv("APP_NAME")
  config.App.Port = toint(os.Getenv("APP_PORT"))
  config.App.Version = os.Getenv("APP_VERSION")
  config.App.Host = os.Getenv("APP_HOST")
  config.App.Base = os.Getenv("APP_BASE")
  config.App.OAuthPort = os.Getenv("APP_OAUTH_PORT")
  config.App.Url = os.Getenv("APP_URL")
  config.App.Secret = os.Getenv("SECRET_KEY")

  config.OAuth.ClientID = os.Getenv("OAUTH_CLIENTID")
  config.OAuth.Secret = os.Getenv("OAUTH_SECRET")
  config.OAuth.Host = os.Getenv("OAUTH_HOST")
  config.OAuth.Web = os.Getenv("OAUTH_WEB")

  config.Metric.Namespace = os.Getenv("METRIC_NAMESPACE ")
  config.Metric.Port = toint(os.Getenv("METRIC_PORT"))

  config.Docs.Host = os.Getenv("DOCS_HOST")
  config.Docs.Path = os.Getenv("DOCS_PATH")
  config.Docs.Port = toint(os.Getenv("DOCS_PORT"))

  config.Remote.Server = os.Getenv("REMOTE_SERVER")
  config.Remote.Port = os.Getenv("REMOTE_PORT")

  return config
}

func toint(value string) int {
  rtn, err := strconv.Atoi(value)
  if err != nil {
    logrus.Fatal(err)
    return 0
  }
  return rtn
}
