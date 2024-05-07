package controllers

import (
	"api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// Controle de requisição
func GetInfo(context *gin.Context) { // Define uma função chamada "getInfo" que recebe um ponteiro para um objeto de contexto Gin
	context.JSON(200, models.Alunos)
}

func GetName(context *gin.Context) {
	// Extrai o parâmetro "name" da URL da requisição e o armazena na variável 'name'.
	name := context.Params.ByName("name")
	context.JSON(200, gin.H{
		"A api diz ": "ola " + name + ", como vai?",
	})
}
