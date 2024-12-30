package models

type Task struct {
    ID         string   `json:"id" bson:"_id,omitempty"` // MongoDB 的 ID
    Title      string   `json:"title" bson:"title"`
    Priority   int      `json:"priority" bson:"priority"`
    Tags       []string `json:"tags" bson:"tags"`
    Attachments []string `json:"attachments" bson:"attachments"`
}
