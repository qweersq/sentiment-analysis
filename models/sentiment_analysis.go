package models

import "time"

type SentimentAnalysis struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SentimentType   string    `gorm:"type:text" json:"sentiment_type"`
	CommentID       uint      `gorm:"not null" json:"-"`
	Comment         Comments  `gorm:"foreignKey:CommentID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"comment_id"`
	ConfidenceLevel uint      `gorm:"not null" json:"confidence_level"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"update_at,omitempty"`
	DeletedAt      time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

func (SentimentAnalysis) TableName() string {
	return "sentiment_analysis"
}
