package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/eniworoeva/JWT-auth/controllers"
)
func AuthRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
	
}