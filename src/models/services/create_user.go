package services

import (
	"fmt"
	errorhandler "gin-gonic-api/src/config/errorHandler"
	"gin-gonic-api/src/config/logger"
	model "gin-gonic-api/src/models"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *errorhandler.RestErr {
	logger.Info("Creating user in the database", zap.String("model", "CreateUser"))

	userDomain.EncryptPassword()

	fmt.Printf("User to be created: %+v\n", userDomain.GetPassword())

	return nil
}
