package server

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/mlambda-net/identity/pkg/domain/utils"
	"log"
	"os"
)

func (s *server) LoadConfig(params ...string) {

	env := os.Getenv("env")
	if env == "" {
		env = "dev"
	}

	if len(params) > 0 {
		env = params[0]
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	s.config = &utils.Configuration{Env: env}

	fileName := fmt.Sprintf("%s/env/%s_server_config.yml", path, env)
	_ = configor.Load(s.config, fileName)

	s.config.Db.Host = os.Getenv("DB_HOST")
	s.config.Db.User = os.Getenv("DB_USER")
	s.config.Db.Password = os.Getenv("DB_PASSWORD")
	s.config.Db.Port = os.Getenv("DB_PORT")
	s.config.Db.Schema = os.Getenv("DB_DATA")

}
