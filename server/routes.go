package server

import (
	"golang-rest-api-starter/controllers"
	"golang-rest-api-starter/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Routes for healthcheck of api server
	healthcheck := router.Group("health")
	{
		health := new(controllers.HealthController)
		ping := new(controllers.PingController)
		healthcheck.GET("/health", health.Status)
		healthcheck.GET("/ping", ping.Ping)
	}

	//Routes for swagger
	swagger := router.Group("swagger")
	{
		// programatically set swagger info
		docs.SwaggerInfo.Title = "Golang REST API Starter"
		docs.SwaggerInfo.Description = "This is a sample backend written in Go."
		docs.SwaggerInfo.Version = "1.0"
		// docs.SwaggerInfo.Host = "cloudfactory.swagger.io"
		// docs.SwaggerInfo.BasePath = "/v1"

		swagger.GET("/*any", ginSwagger.WrapHandler(swagFiles.Handler))
	}

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	return router

}
