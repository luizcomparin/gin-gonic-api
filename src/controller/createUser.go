package controller

import (
	"fmt"
	errorhandler "gin-gonic-api/src/config/errorHandler"
	"gin-gonic-api/src/controller/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	fmt.Println("userRequest: ", userRequest)
	// O ShouldBindJSON transforma a Struct em JSON e valida os campos retornando um erro do tipo BadRequestError
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := errorhandler.NewBadRequestError(fmt.Sprintf("Campos inválidos: %s", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
}
