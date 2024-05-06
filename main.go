package main

import "github.com/gin-gonic/gin"

// Controle de requisição
func getInfo(context *gin.Context) { // Define uma função chamada "getInfo" que recebe um ponteiro para um objeto de contexto Gin
	context.JSON(200, gin.H{ // Chama o método "JSON" do objeto de contexto para enviar uma resposta JSON com status 200 (OK)
		"id":   "1",
		"nome": "kaynan",
	})
}

func main() {
	r := gin.Default()
	r.GET("/server", getInfo)
	r.Run()
}
