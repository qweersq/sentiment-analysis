package models

import (
	"time"
)

type Sentiment struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SentimentType string    `gorm:"type:text" json:"sentiment_type"`
	Created_at    time.Time `json:"created_at,omitempty"`
	Update_at     time.Time `json:"update_at,omitempty"`
	Deleted_at    time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

func (Sentiment) TableName() string {
	return "sentiments"
}
