package main

import (
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func JWTHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "JWT Middleware",
	})
}
func LogHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World 2",
	})
}
func main() {
	r := gin.Default()

	r.GET("", helloHandler)

	r.GET("muitos-handlers-ou-middlewares",
		JWTHandler,
		LogHandler,
		func(c *gin.Context) {
			// implementação do handler
		},
	)

	r.GET("cep/:id", func(c *gin.Context) {
		// Para pegar o parametro "/cep/12345-678"
		cepId := c.Param("id")
		// Para pegar o query "?cep=12345-678"
		cep := c.Query("cep")
		c.JSON(200, gin.H{
			"Parametro Id:": cepId,
			"Query CEP:":    cep,
		})
	})

	err := r.Run(":8787")
	if err != nil {
		return
	}
}
