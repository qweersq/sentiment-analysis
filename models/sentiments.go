package models

import (
	"time"

	"gorm.io/gorm"
)

type Sentiment struct {
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	SentimentType string `gorm:"type:text" json:"sentiment_type"`
	Created_at    time.Time
	Update_at     time.Time
	Deleted_at    gorm.DeletedAt `gorm:"index"`
}

func (Sentiment) TableName() string {
	return "sentiments"
}
