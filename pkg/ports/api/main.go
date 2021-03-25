package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mlambda-net/identity/pkg/infrastructure/api"
	"github.com/mlambda-net/identity/pkg/ports/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	services := api.NewApi()

	docs.SwaggerInfo.Title = "Swagger Identity API"
	docs.SwaggerInfo.Description = "This is the api for the identity service."
	docs.SwaggerInfo.Version = services.GetVersion()
	docs.SwaggerInfo.Host = services.GetHost()
	docs.SwaggerInfo.BasePath = services.Base()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}



	go func() {
		r := gin.New()
		r.GET("/identity/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    _ = r.Run(":8002")
	}()

	services.Start()

}
