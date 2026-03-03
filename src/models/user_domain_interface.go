package model

type UserDomainInterface interface {
	GetEmail() string
	GetName() string
	GetPassword() string
	GetAge() int
	GetID() string

	SetID(string)
	EncryptPassword()
}

func NewUserDomain(name, email, password string, age int) UserDomainInterface {
	return &userDomain{"", email, name, password, age}
}
