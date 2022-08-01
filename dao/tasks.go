package dao

type Task struct {
	SerialID    int64  `json:"serial_id"`
	TaskName    string `json:"task"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	IsCompleted bool   `json:"is_completed"`
	IsDeleted   bool   `json:"is_deleted"`
}