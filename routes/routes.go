package routes

import (
	"backEndGo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/hello", controllers.HelloHandler)
} 