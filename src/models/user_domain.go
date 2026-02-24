package model

import (
	"crypto/md5"
	"encoding/hex"
	errorhandler "gin-gonic-api/src/config/errorHandler"
)

// UserDomain é a estrutura de dados que representa o domínio de usuário na aplicação.
// Ela é utilizada para mapear os dados de entrada (UserRequest) e os dados de saída (UserResponse)
// em um formato que pode ser manipulado internamente pela aplicação.
type UserDomain struct {
	Name     string
	Email    string
	Password string
	Age      int
}

func NewUserDomain(name, email, password string, age int) UserDomainInterface {
	return &UserDomain{email, name, password, age}
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *errorhandler.RestErr
	UpdateUser(string) *errorhandler.RestErr
	FindUser(string) *errorhandler.RestErr
	DeleteUser(string) *errorhandler.RestErr
}
