package api

import (
  "context"
  "fmt"
  "github.com/etherlabsio/healthcheck"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/command"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/conf"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/oauth"
  "github.com/mlambda-net/identity/pkg/infrastructure/api/query"
  "github.com/mlambda-net/net/pkg/metrics"
  "github.com/mlambda-net/net/pkg/net"
  "github.com/sirupsen/logrus"
  "time"
)

type Api interface {
  GetVersion() string
  GetHost() string
  Start()
  Base() string
}

type setup struct {
  config  *conf.Configuration
  command command.Command
  query   query.Query
}

func (s *setup) Base() string {
   if s.config.App.Base != "" {
     return fmt.Sprintf("/%s", s.config.App.Base)
   }
   return ""
}

func NewApi() Api  {
  return &setup{config: conf.LoadConfig()}
}

func (s *setup) GetVersion() string {
  version := s.config.App.Version
  if version == "" {
    version = "0.0.0"
  }
  return version
}

func (s *setup) GetHost() string {
  if s.config.App.Host == "localhost" {
    return fmt.Sprintf("%s:%d", s.config.App.Host, s.config.App.Port)
  }
  return s.config.App.Host
}

func (s *setup) Start() {
  logrus.Infof("Start the API %s on %d", s.config.App.Host, s.config.App.Port)
  logrus.Infof("Connecting to: %s  %s", s.config.Remote.Server, s.config.Remote.Port)
  client := net.NewClient(s.config.Remote.Server, s.config.Remote.Port)

  user := client.Actor("user")
  auth := client.Actor("auth")
  app := client.Actor("app")
  role := client.Actor("role")

  s.command = command.NewCommand(user, auth, app, role, s.config)
  s.query = query.NewQuery(user)

  local := net.NewApi(int32(s.config.App.Port), int32(s.config.Metric.Port))
  local.Metrics(func(mc *metrics.Configuration) {
    mc.App.Name = s.config.App.Name
    mc.App.Env = s.config.Env
    mc.App.Path = "/check/identity"
    mc.App.Version = s.config.App.Version
    mc.Metric.Namespace = s.config.Metric.Namespace
    mc.Metric.SubSystem = "security"
  })
  local.Register(func(r net.Route) {
    s.command.Register(r)
    s.query.Register(r)
  })

  local.Checks(
    healthcheck.WithTimeout(5 * time.Second),
    healthcheck.WithChecker("server",healthcheck.CheckerFunc(func(ctx context.Context) error {
      status, err := client.Health()
      if err != nil{
        return err
      }
      if !status.Success {
        return fmt.Errorf("the api is unhealthy %s",status.Message)
      }
      return nil
    })),
  )

  logrus.Infof("Listen API on %d", s.config.App.Port)
  logrus.Infof("Listen Metrics on %d", s.config.Metric.Port)
  logrus.Infof("Listen Docs on %d", s.config.Docs.Port)
  local.Start()

  server := oauth.NewOAuthServer(s.config, auth)
  server.Start()
  local.Wait()

}
