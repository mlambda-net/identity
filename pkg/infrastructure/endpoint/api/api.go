package api

import (
  "context"
  "fmt"
  "github.com/etherlabsio/healthcheck"
  "github.com/mlambda-net/identity/pkg/infrastructure/endpoint/api/command"
  "github.com/mlambda-net/identity/pkg/infrastructure/endpoint/api/conf"
  "github.com/mlambda-net/identity/pkg/infrastructure/endpoint/api/query"
  "github.com/mlambda-net/net/pkg/metrics"
  "github.com/mlambda-net/net/pkg/net"
  "os"
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
  client  net.Client
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
  version := os.Getenv("VERSION")
  if version == "" {
    version = "0.0.0"
  }
  return version
}

func (s *setup) GetHost() string {
  return fmt.Sprintf("%s:%d",s.config.App.Host, s.config.App.Port)
}

func (s *setup) Start() {
  client := net.NewClient(s.config.Remote.Server, s.config.Remote.Port)
  user := client.Actor("user")
  auth := client.Actor("auth")

  s.command = command.NewCommand(user, auth)
  s.query = query.NewQuery(user)


  local := net.NewApi(s.config.App.Port, s.config.Metric.Port)
  local.Metrics(func(mc *metrics.Configuration) {
    mc.App.Name = s.config.App.Name
    mc.App.Env = s.config.Env
    mc.App.Path = "/check/user"
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
     // client.Health()
      return nil
    })),
  )
  local.Start()

  local.Wait()

}
