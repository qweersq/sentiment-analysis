package repository

import (
	"sentiment/dto"
	"sentiment/models"

	"gorm.io/gorm"
)

type LecturerRepository interface {
	InsertLecturer(lecturers models.Lecturers) models.Lecturers
	UpdateLecturer(lecturers dto.LecturerUpdateDTO) dto.LecturerUpdateDTO
	GetAllLecturer() []models.Lecturers
	GetLecturerByID(id uint64) models.Lecturers
	DeleteLecturer(id uint64) error
}

type lecturerConnection struct {
	connection *gorm.DB
}

func NewLecturerRepository(db *gorm.DB) LecturerRepository {
	return &lecturerConnection{
		connection: db,
	}
}

func (db *lecturerConnection) InsertLecturer(lecturers models.Lecturers) models.Lecturers {
	db.connection.Preload("StudyProgram").Save(&lecturers)
	return lecturers
}

func (db *lecturerConnection) UpdateLecturer(lecturers dto.LecturerUpdateDTO) dto.LecturerUpdateDTO {
	db.connection.Save(&lecturers)
	return lecturers
}

func (db *lecturerConnection) GetAllLecturer() []models.Lecturers {
	var lecturers []models.Lecturers
	db.connection.Preload("StudyProgram").Find(&lecturers)
	return lecturers
}

func (db *lecturerConnection) GetLecturerByID(id uint64) models.Lecturers {
	var lecturer models.Lecturers
	db.connection.Preload("StudyProgram").Find(&lecturer, id)
	return lecturer
}

func (db *lecturerConnection) DeleteLecturer(id uint64) error {
	lecturer := models.Lecturers{}
	result := db.connection.Delete(&lecturer, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
