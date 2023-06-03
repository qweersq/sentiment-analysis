package repository

import (
	"sentiment/models"

	"gorm.io/gorm"
)

type StudyProgramRepository interface {
	InsertStudyProgram(studyProgram models.StudyPrograms) models.StudyPrograms
	UpdateStudyProgram(studyProgram models.StudyPrograms) models.StudyPrograms
	DeleteStudyProgram(id uint64) error
	AllStudyProgram() []models.StudyPrograms
	FindStudyProgramByID(studyProgramID uint64) models.StudyPrograms
	FindStudyProgramByCode(studyProgramName string) models.StudyPrograms
}

type studyProgramConnection struct {
	connection *gorm.DB
}

func NewStudyProgramRepository(db *gorm.DB) StudyProgramRepository {
	return &studyProgramConnection{
		connection: db,
	}
}

func (db *studyProgramConnection) InsertStudyProgram(studyProgram models.StudyPrograms) models.StudyPrograms {
	db.connection.Save(&studyProgram)
	return studyProgram
}

func (db *studyProgramConnection) UpdateStudyProgram(studyProgram models.StudyPrograms) models.StudyPrograms {
	db.connection.Save(&studyProgram)
	return studyProgram
}

func (db *studyProgramConnection) DeleteStudyProgram(id uint64) error {
	studyProgram := models.StudyPrograms{}
	result := db.connection.Delete(&studyProgram, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *studyProgramConnection) AllStudyProgram() []models.StudyPrograms {
	var studyPrograms []models.StudyPrograms
	db.connection.Find(&studyPrograms)
	return studyPrograms
}

func (db *studyProgramConnection) FindStudyProgramByID(studyProgramID uint64) models.StudyPrograms {
	var studyProgram models.StudyPrograms
	db.connection.Find(&studyProgram, studyProgramID)
	return studyProgram
}

func (db *studyProgramConnection) FindStudyProgramByCode(studyProgramCode string) models.StudyPrograms {
	var studyProgram models.StudyPrograms
	db.connection.Where("code = ?", studyProgramCode).Take(&studyProgram)
	return studyProgram
}
