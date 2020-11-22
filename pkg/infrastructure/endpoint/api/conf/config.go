package conf

import (
	"fmt"
	"github.com/jinzhu/configor"
	"log"
	"os"
)

type Configuration struct {
	Env string `default:"dev"`

	App struct {
		Name    string
		Version string `default:"1.0.0"`
		Port    int32 `default:"8080"`
		Host    string `default:"localhost"`
		Base    string `default:""`
	}

	Metric struct {
		Port      int32 `default:"8081"`
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

func LoadConfig(params ...string) *Configuration {

	env := os.Getenv("ENV")
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

	config := &Configuration{Env: env}

	fileName := fmt.Sprintf("%s/env/%s_api_config.yml", path, env)
	_ = configor.Load(config, fileName)
	return config
}
