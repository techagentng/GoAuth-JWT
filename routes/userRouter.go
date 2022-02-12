package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/techagentng/GOAUTH-JJWT/controllers"
	"github.com/techagentng/GOAUTH-JJWT/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUserID())
}
