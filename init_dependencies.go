package main

import (
	"gin-gonic-api/src/controller"
	"gin-gonic-api/src/models/repository"
	"gin-gonic-api/src/models/services"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := services.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
