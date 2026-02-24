package model

import (
	"fmt"
	errorhandler "gin-gonic-api/src/config/errorHandler"
	"gin-gonic-api/src/config/logger"

	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *errorhandler.RestErr {
	logger.Info("Creating user in the database", zap.String("model", "CreateUser"))

	ud.EncryptPassword()

	fmt.Println("User to be created: %+v\n", ud)

	return nil
}
