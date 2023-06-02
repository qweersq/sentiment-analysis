package models

import "time"

type StudyPrograms struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudyProgramCode string    `gorm:"type:text" json:"study_program_code"`
	StudyProgramName string    `gorm:"type:text" json:"study_program_name"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"update_at,omitempty"`
	DeletedAt        time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

func (StudyPrograms) TableName() string {
	return "study_programs"
}
