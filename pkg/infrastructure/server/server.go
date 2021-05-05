package server

import (
	"fmt"
  "github.com/mlambda-net/identity/pkg/application/app"
  "github.com/mlambda-net/identity/pkg/application/auth"
  "github.com/mlambda-net/identity/pkg/application/roles"
  "github.com/mlambda-net/identity/pkg/application/user"
  "github.com/mlambda-net/identity/pkg/domain/utils"
  "github.com/mlambda-net/identity/pkg/infrastructure/db"
  "github.com/mlambda-net/net/pkg/remote"
	log "github.com/sirupsen/logrus"
  "os"
  "sync"
)

type Server interface {
	Start()
	Stop()
	Wait()
}

type server struct {
	wg     *sync.WaitGroup
	config *utils.Configuration
	remote remote.Server
}

func (s server) Wait() {
	s.wg.Wait()
}

func NewServer() Server {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	return server{wg: wg}
}

func (s server) Start() {

  log.SetFormatter(&log.TextFormatter{
    FullTimestamp: true,
  })
  log.SetOutput(os.Stdout)


  defer func() {
    if err := recover(); err != nil {
      log.Println("panic occurred:", err)
    }
  }()

  s.LoadConfig()
  s.remote = remote.NewServer()
  go func() {
    log.Info(fmt.Sprintf("starting the server %s in the port %s", s.config.App.Name, s.config.App.Port))
    s.remote.Register("user", user.NewUserProps(s.config), false, []string{"identity_admin"})
    s.remote.Register("auth", auth.NewAuthProps(s.config), false, []string{})
    s.remote.Register("app", app.NewAppActor(s.config), false, []string{"identity_admin"})
    s.remote.Register("role", roles.NewRolesProps(s.config), false,[]string{"identity_admin"})

    s.remote.Check(func(status *remote.Status) {
      err := db.AliveDB(s.config)
      if err != nil {
        status.Add(false, "database", err.Error())
        log.Errorf("The database  is not available. %s", err.Error())
      } else {
        status.Add(true, "database", "ok")
      }

      err = db.AliveCache(s.config)
      if err != nil {
        status.Add(false, "cache", err.Error())
        log.Errorf("The cache is not available. %s", err.Error())
      } else {
        status.Add(true, "cache", "ok")
      }
    })

    s.remote.Start(fmt.Sprintf(":%s", s.config.App.Port))

    err := db.AliveDB(s.config)
    if err != nil {
      log.Errorf("The database  is not available. %s", err.Error())
    }

    err  = db.AliveCache(s.config)
    if err != nil {
      log.Errorf("The cache is not available. %s", err.Error())
    }

    s.wg.Wait()
  }()
}

func (s server) Stop() {
	s.remote.Stop()
	s.wg.Done()
}
