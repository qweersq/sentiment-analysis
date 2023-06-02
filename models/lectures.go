package models

import "time"

type Lectures struct {
	ID             uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Lecturer_Code  string        `gorm:"type:text" json:"lecturer_code"`
	Lecturer_Name  string        `gorm:"type:text" json:"lecturer_name"`
	StudyProgramID uint          `gorm:"not null" json:"-"`
	StudyProgram   StudyPrograms `gorm:"foreignKey:StudyProgramID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"study_program_id"`
	CreatedAt     time.Time     `json:"created_at,omitempty"`
	UpdatedAt      time.Time     `json:"update_at,omitempty"`
	DeletedAt     time.Time     `gorm:"index"  json:"deleted_at,omitempty"`
}

func (Lectures) TableName() string {
	return "lectures"
}
