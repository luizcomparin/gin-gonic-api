package view

import (
	"gin-gonic-api/src/controller/model/response"
	model "gin-gonic-api/src/models"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}
