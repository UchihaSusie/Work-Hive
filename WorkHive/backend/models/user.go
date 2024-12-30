package models

type User struct {
    Username         string `json:"username"`
    Password         string `json:"password"`
    Role             string `json:"role"`
    LoginCount       int    `json:"login_count"` // 登录次数
    CompletedTasks   int    `json:"completed_tasks"` // 完成的任务数量
    ProductivityScore float64 `json:"productivity_score"` // 生产力评分
}

func (u *User) CalculateProductivityScore() {
    if u.LoginCount == 0 {
        u.ProductivityScore = 0
    } else {
        u.ProductivityScore = float64(u.CompletedTasks) / float64(u.LoginCount) * 100 // 计算生产力评分
    }
}