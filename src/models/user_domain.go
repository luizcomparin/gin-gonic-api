package model

// userDomain é a estrutura de dados que representa o domínio de usuário na aplicação.
// Ela é utilizada para mapear os dados de entrada (UserRequest) e os dados de saída (UserResponse)
// em um formato que pode ser manipulado internamente pela aplicação.
// Structs e propriedades começando com letra MINUSCULA são privadas.
type userDomain struct {
	id       string
	name     string
	email    string
	password string
	age      int
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
func (ud *userDomain) GetID() string {
	return ud.id
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}
