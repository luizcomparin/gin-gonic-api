package controller

import "github.com/gin-gonic/gin"

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	diabo := c.Param("id")

	c.JSON(200, gin.H{
		"id": diabo,
	})
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {

}
