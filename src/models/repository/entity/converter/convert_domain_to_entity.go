package converter

import (
	model "gin-gonic-api/src/models"
	"gin-gonic-api/src/models/repository/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Name:     domain.GetName(),
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Age:      domain.GetAge(),
	}
}
