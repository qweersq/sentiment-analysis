package models

import "time"

type Comments struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Comment    string    `gorm:"type:text" json:"comment"`
	CourseID   uint      `gorm:"not null" json:"-"`
	Course     Courses   `gorm:"foreignKey:CourseID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"course_id"`
	LecturerID uint      `gorm:"not null" json:"-"`
	Lecturer   Lectures  `gorm:"foreignKey:LecturerID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"lecturer_id"`
	SchoolYear uint      `gorm:"not null" json:"school_year"`
	Semester   uint      `gorm:"not null" json:"semester"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"update_at,omitempty"`
	DeletedAt  time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

func (Comments) TableName() string {
	return "comments"
}
