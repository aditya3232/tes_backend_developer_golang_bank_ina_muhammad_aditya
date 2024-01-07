package users

import "time"

type User struct {
	ID        int        `gorm:"primary_key;column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	Email     string     `gorm:"unique;column:email" json:"email"`
	Password  string     `gorm:"column:password" json:"password"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTTime" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
