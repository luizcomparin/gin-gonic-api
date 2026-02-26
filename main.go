package main

import (
	"fmt"
	"gin-gonic-api/src/config/logger"
	"gin-gonic-api/src/controller"
	"gin-gonic-api/src/controller/routes"
	"gin-gonic-api/src/models/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	logger.Info("Iniciando o servidor...")

	errdotenv := godotenv.Load()
	if errdotenv != nil {
		log.Fatal("Erro ao carregar o arquivo .env: ", errdotenv)
		return
	}
	fmt.Printf("Env TEST: %s\n", os.Getenv("TEST"))

	// Initi dependency injection
	service := services.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	routes.InitRoutes(&r.RouterGroup, userController)

	err := r.Run(":8787")
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
