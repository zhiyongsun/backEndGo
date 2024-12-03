package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	dsn := "admin1:Szy785037@tcp(rm-bp1c60m16vethpik91o.mysql.rds.aliyuncs.com:3306)/hefipal?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 初始化生成器
	g := gen.NewGenerator(gen.Config{
		OutPath: "./models/gen", // 生成代码的输出目录
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式配置
	})

	// 使用数据库连接
	g.UseDB(db)

	// 生成所有表的模型
	g.GenerateAllTable()

	// 执行生成
	g.Execute()
} 