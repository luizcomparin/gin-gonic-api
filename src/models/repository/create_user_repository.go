package repository

import (
	"context"
	errorhandler "gin-gonic-api/src/config/errorHandler"
	"gin-gonic-api/src/config/logger"
	model "gin-gonic-api/src/models"
	"gin-gonic-api/src/models/repository/entity/converter"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

const (
	MONGODB_COLLECTION_NAME = "MONGODB_COLLECTION_NAME"
)

// var (MONGODB_COLLECTION_NAME = os.Getenv("MONGODB_COLLECTION_NAME"))

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *errorhandler.RestErr) {

	logger.Info("Criando usuário no banco de dados...", zap.String("repository", "CreateUser"))
	collection_name := os.Getenv(MONGODB_COLLECTION_NAME)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error inserting user on collection", err, zap.String("repository", "CreateUser"))
		return nil, errorhandler.NewInternalServerError("Erro ao inserir user na collection: " + err.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)

	logger.Info(
		"User created successfully",
		zap.String("repository", "CreateUser"),
		zap.String("userID", value.Id.Hex()),
	)

	return converter.ConvertEntityToDomain(*value), nil
}
