package main

import (
	"context"
	"gin-gonic-api/src/config/database/mongodb"
	"gin-gonic-api/src/config/logger"
	"gin-gonic-api/src/controller/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	APP_PORT = "APP_PORT"
)

func main() {
	r := gin.Default()
	logger.Info("Iniciando o servidor...")

	errdotenv := godotenv.Load()
	if errdotenv != nil {
		log.Fatal("Erro ao carregar o arquivo .env: ", errdotenv)
		return
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatal("Erro ao conectar ao MongoDB: ", err)
		return
	}

	userController := initDependencies(database)

	routes.InitRoutes(&r.RouterGroup, userController)

	if err := r.Run(":" + os.Getenv(APP_PORT)); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
