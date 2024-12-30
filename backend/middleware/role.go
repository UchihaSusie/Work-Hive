package middleware

import (
	"github.com/gin-gonic/gin"
	"backend/database"
)

func RoleRequired(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.Request.Header.Get("Authorization") // 从请求头获取会话 ID
		if sessionID == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 从 Redis 获取用户角色
		userRole, err := database.GetRedisClient().HGet(database.Ctx, sessionID, "role").Result()
		if err != nil || userRole != role {
			c.JSON(403, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
