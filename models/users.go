package models

import "time"

type User struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Username   string `gorm:"type:varchar(255)" json:"username"`
	Email      string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password   string `gorm:"->;<-;not null" json:"-"`
	Created_at time.Time
	Update_at  time.Time
	Deleted_at time.Time `gorm:"index"`
	Token      string    `gorm:"-" json:"token,omitempty"`
}

func (User) TableName() string {
	return "users"
}
