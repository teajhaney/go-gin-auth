package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/user"

	"github.com/gin-gonic/gin"
)


func NewRouter (app *app.App) *gin.Engine{
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Register health routes
	router.GET("/health", health)

	// Register auth routes
  user.RegisterUserRoutes(router, app)
	return router
}
