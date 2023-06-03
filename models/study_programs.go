package models

import "time"

type StudyPrograms struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code      string    `gorm:"type:text" json:"code"`
	Name      string    `gorm:"type:text" json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
	DeletedAt time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

func (StudyPrograms) TableName() string {
	return "study_programs"
}
