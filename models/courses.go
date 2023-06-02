package models

import "time"

type Courses struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Course_Code string    `gorm:"type:text" json:"course_code"`
	Course_Name string    `gorm:"type:text" json:"course_name"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"update_at,omitempty"`
	DeletedAt  time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

func (Courses) TableName() string {
	return "courses"
}
