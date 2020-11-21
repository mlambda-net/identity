package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mlambda-net/identity/pkg/infrastructure/endpoint/api"
	"github.com/mlambda-net/identity/pkg/infrastructure/ports/api/docs"
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
		r.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.Run(":8003")
	}()

	services.Start()

}
