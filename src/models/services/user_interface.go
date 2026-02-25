package services

import (
	errorhandler "gin-gonic-api/src/config/errorHandler"
	model "gin-gonic-api/src/models"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *errorhandler.RestErr
	UpdateUser(string, model.UserDomainInterface) *errorhandler.RestErr
	FindUser(string) (*model.UserDomainInterface, *errorhandler.RestErr)
	DeleteUser(string) *errorhandler.RestErr
}
