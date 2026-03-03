package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// Entity para o MongoDB, mapeando os campos da aplicação para a estrutura do banco de dados .
type UserEntity struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Age      int                `bson:"age"`
}
