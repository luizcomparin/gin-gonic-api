package services

import (
	errorhandler "gin-gonic-api/src/config/errorHandler"
	model "gin-gonic-api/src/models"
	"gin-gonic-api/src/models/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *errorhandler.RestErr)
	UpdateUser(string, model.UserDomainInterface) *errorhandler.RestErr
	FindUser(string) (*model.UserDomainInterface, *errorhandler.RestErr)
	DeleteUser(string) *errorhandler.RestErr
}
