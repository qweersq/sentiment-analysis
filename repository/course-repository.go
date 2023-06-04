package repository

import (
	"sentiment/models"

	"gorm.io/gorm"
)

type CourseRepository interface {
	InsertCourse(courses models.Courses) models.Courses
	UpdateCourse(courses models.Courses) models.Courses
	GetAllCourse() []models.Courses
	GetCourseByID(id uint64) models.Courses
	DeleteCourse(id uint64) error
}

type courseConnection struct {
	connection *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseConnection{
		connection: db,
	}
}

func (db *courseConnection) InsertCourse(courses models.Courses) models.Courses {
	db.connection.Save(&courses)
	return courses
}

func (db *courseConnection) UpdateCourse(courses models.Courses) models.Courses {
	db.connection.Save(&courses)
	return courses
}

func (db *courseConnection) GetAllCourse() []models.Courses {
	var courses []models.Courses
	db.connection.Preload("StudyProgram").Find(&courses)
	return courses
}

func (db *courseConnection) GetCourseByID(id uint64) models.Courses {
	var course models.Courses
	db.connection.Preload("StudyProgram").Find(&course, id)
	return course
}

func (db *courseConnection) DeleteCourse(id uint64) error {
	course := models.Courses{}
	result := db.connection.Delete(&course, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
