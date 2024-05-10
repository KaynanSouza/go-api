package controllers

import (
	"api-go-gin/database"
	"api-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controle de requisição
func GetInfo(context *gin.Context) { // Define uma função chamada "getInfo" que recebe um ponteiro para um objeto de contexto Gin
	var alunos []models.Aluno
	database.DB.Find(&alunos) // Buscar todos os alunos dentro do banco de dados
	context.JSON(200, alunos)
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

func FindId(context *gin.Context) {
	var aluno models.Aluno
	id := context.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"not found": "error 404",
		})
		return
	}

	context.JSON(http.StatusOK, aluno)
}

func DeleteAluno(context *gin.Context) {
	var aluno models.Aluno
	id := context.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Not found": "Error 404",
		})
		return
	}

	database.DB.Delete(&aluno, id)
	context.JSON(http.StatusOK, gin.H{
		"Sucefull": "Aluno deletado",
	})
}

func UpdateAluno(context *gin.Context) {
	var aluno models.Aluno
	id := context.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := context.ShouldBindJSON(&aluno); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	context.JSON(http.StatusOK, gin.H{
		"Sucefull": "Aluno modificado",
	})

}

func FindCpf(context *gin.Context) {
	var aluno models.Aluno
	cpf := context.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}, cpf).First(&aluno)

	if aluno.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"Not found": "Error 404",
		})
		return
	}

	context.JSON(http.StatusOK, aluno)

}
