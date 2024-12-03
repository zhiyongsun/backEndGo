package controllers

import (
	"backEndGo/models/gen" // 导入生成的模型
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	// 这里可以使用生成的模型和查询方法
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
} 