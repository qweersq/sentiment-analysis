package models

import "time"

type Lecturers struct {
	ID             uint64        `gorm:"primaryKey;autoIncrement" json:"id"`
	Code           string        `gorm:"type:text" json:"code"`
	Name           string        `gorm:"type:text" json:"name"`
	StudyProgramID uint64        `gorm:"not null" json:"-"`
	StudyProgram   StudyPrograms `gorm:"foreignKey:StudyProgramID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"study_program_id"`
	CreatedAt      time.Time     `json:"created_at,omitempty"`
	UpdatedAt      time.Time     `json:"update_at,omitempty"`
	DeletedAt      time.Time     `gorm:"index"  json:"deleted_at,omitempty"`
}

func (Lecturers) TableName() string {
	return "lecturers"
}
