package repository

import (
	errorhandler "gin-gonic-api/src/config/errorHandler"
	model "gin-gonic-api/src/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *errorhandler.RestErr)
}
