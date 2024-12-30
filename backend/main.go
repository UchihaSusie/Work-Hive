package main

import (
	"backend/database"
	"backend/middleware"
	"backend/routes"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 初始化 Redis 和 MongoDB
	database.InitRedis()
	database.InitMongoDB()

	r := gin.Default()

	// 路由设置
	r.POST("/api/login", routes.Login)
	r.POST("/api/tasks", middleware.RoleRequired("admin"), routes.CreateTask)
	r.GET("/api/tasks", routes.GetTasks)
	r.GET("/api/user/productivity", routes.GetUserProductivity)

	r.Run() // 默认在 8080 端口运行
}