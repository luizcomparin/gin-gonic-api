package mongodb

import (
	"context"
	"crypto/tls"

	// "fmt"
	"gin-gonic-api/src/config/logger"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Essa seria a variável global que armazenará a instância do cliente MongoDB,
// permitindo que seja reutilizada em toda a aplicação sem a necessidade de criar
// múltiplas conexões e sem gerar efeito cascata de dependências e construtores.
// var mongoDBClient *mongo.Client

var (
	MONGODB_URI           = "MONGODB_URI"
	MONGODB_DATABASE_NAME = "MONGODB_DATABASE_NAME"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	MONGO_URI := os.Getenv(MONGODB_URI)
	MONGO_DATABASE := os.Getenv(MONGODB_DATABASE_NAME)
	if MONGO_URI == "" {
		panic(MONGODB_URI + " nao definido no .env")
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	clientOptions := options.Client().
		ApplyURI(MONGO_URI).
		SetServerSelectionTimeout(15 * time.Second).
		SetTLSConfig(&tls.Config{MinVersion: tls.VersionTLS12})

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		// panic(fmt.Errorf("Erro ao iniciar cliente MongoDB: %w", err))
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		// panic(fmt.Errorf("Erro ao pingar MongoDB (verifique URI, TLS e Atlas Network Access): %w", err))
		return nil, err
	}

	// mongoDBClient = client
	logger.Info("Conexao com o MongoDB estabelecida com sucesso!")
	database := client.Database(MONGO_DATABASE)

	return database, nil
}
