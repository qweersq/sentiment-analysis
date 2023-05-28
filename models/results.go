package models

import (
	"time"
)

type Result struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint64    `gorm:"not null" json:"-"`
	User        User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user_id"`
	DataTextID  uint      `gorm:"not null" json:"-"`
	DataText    DataText  `gorm:"foreignkey:DataTextID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"data_text_id"`
	SentimentID uint      `gorm:"not null" json:"-"`
	Sentiment   Sentiment `gorm:"foreignkey:SentimentID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"sentiment_id"`
	Confidence  float64   `gorm:"type:decimal(10,2)" json:"confidence"`
	Created_at  time.Time `json:"created_at,omitempty"`
	Update_at   time.Time `json:"update_at,omitempty"`
	Deleted_at  time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

func (Result) TableName() string {
	return "results"
}
