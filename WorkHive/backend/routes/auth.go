package routes

import (
	"backend/database"
	"backend/models"
	"sync"

	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, models.Response{Status: "error", Message: "Invalid input"})
		return
	}

	var wg sync.WaitGroup
	var role string
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		role = getUserRole(user.Username)
	}()

	wg.Wait()

	if user.Password != "expected_password" {
		c.JSON(401, models.Response{Status: "error", Message: "Invalid credentials"})
		return
	}

	updateUserLoginCount(user.Username)

	sessionID := "some_generated_session_id"

	err = database.GetRedisClient().HSet(database.Ctx, sessionID, "username", user.Username, "role", role).Err()
	if err != nil {
		c.JSON(500, models.Response{Status: "error", Message: "Failed to set session"})
		return
	}

	c.JSON(200, models.Response{Status: "success", Message: "Login successful", Data: gin.H{"session_id": sessionID, "role": role}})
}

func updateUserLoginCount(username string) {
	collection := database.GetMongoClient().Database("your_database").Collection("users")
	_, err := collection.UpdateOne(database.Ctx, bson.M{"username": username}, bson.M{"$inc": bson.M{"login_count": 1}})
	if err != nil {
		log.Printf("Failed to update login count for user %s: %v", username, err)
	}
}

func GetUserProductivity(c *gin.Context) {
	username := c.Request.Header.Get("Username") // 从请求头获取用户名

	var user models.User
	err := database.GetMongoClient().Database("your_database").Collection("users").FindOne(database.Ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		c.JSON(404, models.Response{Status: "error", Message: "User not found"})
		return
	}

	user.CalculateProductivityScore()

	c.JSON(200, models.Response{Status: "success", Message: "User productivity data", Data: user})
}
