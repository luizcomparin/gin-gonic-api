package model

import (
	"crypto/md5"
	"encoding/hex"
)

// userDomain é a estrutura de dados que representa o domínio de usuário na aplicação.
// Ela é utilizada para mapear os dados de entrada (UserRequest) e os dados de saída (UserResponse)
// em um formato que pode ser manipulado internamente pela aplicação.
// Structs e propriedades começando com letra MINUSCULA são privadas.
type userDomain struct {
	name     string
	email    string
	password string
	age      int
}

type UserDomainInterface interface {
	GetEmail() string
	GetName() string
	GetPassword() string
	GetAge() int

	EncryptPassword()
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetAge() int {
	return ud.age
}

func NewUserDomain(name, email, password string, age int) UserDomainInterface {
	return &userDomain{email, name, password, age}
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
