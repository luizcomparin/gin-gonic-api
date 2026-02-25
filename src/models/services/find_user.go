package services

import (
	errorhandler "gin-gonic-api/src/config/errorHandler"
	model "gin-gonic-api/src/models"
)

func (*userDomainService) FindUser(string) (*model.UserDomainInterface, *errorhandler.RestErr) {
	return nil, nil
}
