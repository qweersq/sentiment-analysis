package models

import "gorm.io/gorm"

type Result struct {
	ID          uint `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint `gorm:"not null" json:"user_id"`
	DataTextID  uint `gorm:"not null" json:"data_text_id"`
	SentimentID uint `gorm:"not null" json:"sentiment_id"`
	Confidence  float64 `gorm:"type:decimal(10,2)" json:"confidence"`
}

func (Result) TableName() string {
	return "results"
}
