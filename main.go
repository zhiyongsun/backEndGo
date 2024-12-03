package main

import (
	"backEndGo/config"
	"backEndGo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	config.InitDB()

	// 创建 Gin 引擎
	r := gin.Default()

	// 设置路由
	routes.SetupRoutes(r)

	// 启动服务器
	r.Run(":8080")
} 