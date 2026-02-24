package request

type UserRequest struct {
	// Info sobre validação: https://github.com/go-playground/validator
	// O campo "binding" é utilizado para validar os campos obrigatórios, nesse caso o campo "name" é obrigatório,
	// caso ele não seja preenchido, o ShouldBindJSON irá retornar um erro do tipo BadRequestError
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100,containsany=!@#$%^&*()_+-=~"`
	Age      int    `json:"age" binding:"required,min=1,max=200"`
}
