package converter

import (
	model "gin-gonic-api/src/models"
	"gin-gonic-api/src/models/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Name,
		entity.Password,
		entity.Age,
	)

	domain.SetID(entity.Id.Hex())

	return domain
}
