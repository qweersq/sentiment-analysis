package models

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string    `gorm:"->;<-;not null" json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
	DeletedAt time.Time `gorm:"index"  json:"deleted_at,omitempty"`
	Token     string    `gorm:"-" json:"token,omitempty"`
}

func (User) TableName() string {
	return "users"
}
