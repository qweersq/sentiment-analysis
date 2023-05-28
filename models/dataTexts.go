package models

import (
	"time"

	"gorm.io/gorm"
)

type DataText struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint   `gorm:"not null" json:"user_id"`
	Text       string `gorm:"type:text" json:"text"`
	Created_at time.Time
	Update_at  time.Time
	Deleted_at gorm.DeletedAt `gorm:"index"`
}

func (DataText) TableName() string {
	return "data_texts"
}
