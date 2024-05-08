package controllers

import (
	"api-go-gin/database"
	"api-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
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

func PostAluno(context *gin.Context) {
	var aluno models.Aluno

	if err := context.ShouldBindJSON(&aluno); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&aluno)
	context.JSON(http.StatusOK, aluno)

}
