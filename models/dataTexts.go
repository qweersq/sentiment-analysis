package models

import (
	"time"
)

type DataText struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint64    `gorm:"not null" json:"-"`
	User       User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user_id"`
	Text       string    `gorm:"type:text" json:"text"`
	Created_at time.Time `json:"created_at,omitempty"`
	Update_at  time.Time `json:"update_at,omitempty"`
	Deleted_at time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

func (DataText) TableName() string {
	return "data_texts"
}
