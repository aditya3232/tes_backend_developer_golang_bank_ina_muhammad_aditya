package tasks

import "time"

type Task struct {
	ID          int        `gorm:"primary_key;column:id" json:"id"`
	UserID      int        `gorm:"column:user_id" json:"user_id"`
	Title       string     `gorm:"column:title" json:"title"`
	Description string     `gorm:"column:description" json:"description"`
	Status      string     `gorm:"column:status;default:pending" json:"status"`
	CreatedAt   *time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTTime" json:"updated_at"`
}

func (Task) TableName() string {
	return "tasks"
}
