package services

import (
	errorhandler "gin-gonic-api/src/config/errorHandler"
	model "gin-gonic-api/src/models"
)

func (*userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *errorhandler.RestErr {
	return nil
}
