package services

import (
	errorhandler "gin-gonic-api/src/config/errorHandler"
	"gin-gonic-api/src/config/logger"
	model "gin-gonic-api/src/models"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *errorhandler.RestErr) {
	logger.Info("Creating user in the database", zap.String("service", "CreateUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("service", "CreateUser"))
		return nil, err
	}

	logger.Info(
		"User created successfully",
		zap.String("service", "CreateUser"),
		zap.String("userID", userDomainRepository.GetID()),
	)

	return userDomainRepository, nil
}
