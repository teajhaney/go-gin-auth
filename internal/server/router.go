package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/middleware"
	"go-auth/internal/user"
	"net/http"

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

  //route group
  api := router.Group("/api")

  //auth middleware

  api.Use(middleware.AuthRequired(string(app.Config.JWTSecret)))
 api.GET("/files", func(c *gin.Context){
	userID, ok := middleware.GetUserID(c)
	if !ok{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"userID": userID,
		"files": []any{},
	})
 })


// auth and admin middleware
 api.GET("/products", middleware.RequiredAdmin(), func(c *gin.Context){
	userID, ok := middleware.GetUserID(c)
	if !ok{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	role, ok := middleware.GetRole(c)
	if !ok{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Only admin can access this route",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"userID": userID,
		"role": role,
		"products": []any{},
	})

 })


 //admin middleware used on a group with authorisation 
 admin := api.Group("/admin")
 admin.Use(middleware.RequiredAdmin())
 admin.GET("/dashboard",  func(c *gin.Context){
	userID, ok := middleware.GetUserID(c)
	if !ok{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	role, ok := middleware.GetRole(c)
	if !ok{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Only admin can access this route",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"userID": userID,
		"role": role,
		"dashboard": []any{},
	})
 })
	return router
}
