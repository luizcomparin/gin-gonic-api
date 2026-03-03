package controller

import (
	// "fmt"
	// errorhandler "gin-gonic-api/src/config/errorHandler"
	"gin-gonic-api/src/config/logger"
	"gin-gonic-api/src/config/validation"
	"gin-gonic-api/src/controller/model/request"
	model "gin-gonic-api/src/models"
	"gin-gonic-api/src/view"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	logger.Info("Init CreateUser controller", zap.String("controller", "CreateUser"))
	// O ShouldBindJSON transforma a Struct em JSON e valida os campos retornando um erro do tipo BadRequestError
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error validating user input: ", err, zap.String("controller", "CreateUser"))

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Name, userRequest.Email, userRequest.Password, userRequest.Age)

	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		logger.Error("Error creating user: ", err, zap.String("controller", "CreateUser"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"User created successfully",
		zap.String("controller", "CreateUser"),
		zap.String("userID", domainResult.GetID()),
	)

	createResponse := view.ConvertDomainToResponse(domainResult)

	logger.Info(
		"User created successfully",
		zap.String("controller", "CreateUser"),
	)

	c.JSON(http.StatusOK, createResponse)
}
