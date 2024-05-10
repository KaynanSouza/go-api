package routes

import (
	"api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.GetInfo)
	r.POST("/alunos", controllers.PostAluno)
	r.GET("/alunos/:id", controllers.FindId)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.PATCH("/alunos/:id", controllers.UpdateAluno)
	r.GET("/:name", controllers.GetName)
	r.GET("/alunos/cpf/:cpf", controllers.FindCpf)
	r.Run()
}
