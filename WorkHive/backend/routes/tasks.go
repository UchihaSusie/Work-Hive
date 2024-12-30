package routes

import (
	"backend/database"
	"backend/models"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// 创建任务
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, models.Response{Status: "error", Message: "Invalid input"})
		return
	}

	collection := database.GetMongoClient().Database("your_database").Collection("tasks")
	_, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		c.JSON(500, models.Response{Status: "error", Message: "Failed to create task"})
		return
	}

	c.JSON(201, models.Response{Status: "success", Message: "Task created", Data: task})
}

// 获取所有任务
func GetTasks(c *gin.Context) {
	collection := database.GetMongoClient().Database("your_database").Collection("tasks")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(500, models.Response{Status: "error", Message: "Failed to fetch tasks"})
		return
	}
	defer cursor.Close(context.Background())

	var tasks []models.Task
	if err = cursor.All(context.Background(), &tasks); err != nil {
		c.JSON(500, models.Response{Status: "error", Message: "Failed to decode tasks"})
		return
	}

	c.JSON(200, models.Response{Status: "success", Message: "Tasks fetched", Data: tasks})
}

// 更新任务
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	collection := database.GetMongoClient().Database("your_database").Collection("tasks")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": task})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(200, task)
}

// 删除任务
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	collection := database.GetMongoClient().Database("your_database").Collection("tasks")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(204, nil)
}

func CompleteTask(c *gin.Context) {
	id := c.Param("id")
	username := c.Request.Header.Get("Username") // 从请求头获取用户名

	// 更新任务状态为完成
	collection := database.GetMongoClient().Database("your_database").Collection("tasks")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"completed": true}})
	if err != nil {
		c.JSON(500, models.Response{Status: "error", Message: "Failed to complete task"})
		return
	}

	// 更新用户完成的任务数量
	updateUserCompletedTasks(username)

	c.JSON(200, models.Response{Status: "success", Message: "Task completed"})
}

// 更新用户完成的任务数量的函数
func updateUserCompletedTasks(username string) {
	collection := database.GetMongoClient().Database("your_database").Collection("users")
	_, err := collection.UpdateOne(database.Ctx, bson.M{"username": username}, bson.M{"$inc": bson.M{"completed_tasks": 1}})
	if err != nil {
		log.Printf("Failed to update completed tasks for user %s: %v", username, err)
	}
}
