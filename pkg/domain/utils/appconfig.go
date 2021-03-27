package utils



type AppConfig struct {

  Cache struct {
    Server   string `default:"localhost"`
    Port     int    `default:"6379"`
    Password string `default:""`
    DB       int    `default:"0"`
  }



}
