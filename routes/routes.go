package routes

import (
	"api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/server", controllers.GetInfo)
	r.GET("/:name", controllers.GetName)
	r.Run()
}
