// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"charitybay/internal/usecase"
	"charitybay/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// Swagger docs.
)

// NewRouter -.
// Swagger spec:
// @title		CharityBay API
// @description This is the CharityBay API server.
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.LoggerI, t usecase.Greet) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Routers
	h := handler.Group("/v1")
	{
		newGreetingRoutes(h, t, l)
	}
}
