package routes

import (
	"gin-gonic-api/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("userById/:id", userController.FindUserByID)
	r.GET("userByEmail/:email", userController.FindUserByEmail)
	r.POST("user", userController.CreateUser)
	r.PUT("user", userController.UpdateUser)
	r.DELETE("user", userController.DeleteUser)
}
