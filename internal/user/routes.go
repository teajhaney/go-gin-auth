package user

import (
	"go-auth/internal/app"


	"github.com/gin-gonic/gin"

)



func RegisterUserRoutes(router *gin.Engine, app *app.App	) {

	//create a new note repository and handler
	repo := NewRepo(app.DB)
	service := NewService(repo, string(app.Config.JWTSecret))
	handler := NewHandler(service)


	//group routes under /users
	userGroup := router.Group("/users")
	{
	  
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)

	}

}
