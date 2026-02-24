package routes

import (
	"gin-gonic-api/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("userById/:id", controller.FindUserByID)
	r.GET("userByEmail/:email", controller.FindUserByEmail)
	r.POST("user", controller.CreateUser)
	r.PUT("user", controller.UpdateUser)
	r.DELETE("user", controller.DeleteUser)
}
