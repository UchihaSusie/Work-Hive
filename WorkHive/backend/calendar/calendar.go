package models

import "time"

type Task struct {
    ID          string    `json:"id" bson:"_id,omitempty"` // MongoDB 的 ID
    Title       string    `json:"title" bson:"title"`
    Priority    int       `json:"priority" bson:"priority"`
    Tags        []string  `json:"tags" bson:"tags"`
    Attachments []string  `json:"attachments" bson:"attachments"`
    StartTime   time.Time `json:"start_time" bson:"start_time"` // 任务开始时间
    EndTime     time.Time `json:"end_time" bson:"end_time"`     // 任务结束时间
}
