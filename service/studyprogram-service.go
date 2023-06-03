package service

import (
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"
)

type StudyProgramService interface {
	InsertProdi(studyProgram dto.StudyProgramCreateDTO) models.StudyPrograms
	UpdateProdi(studyProgram dto.StudyProgramUpdateDTO) models.StudyPrograms
	DeleteProdi(id uint64) error
	AllProdi() []models.StudyPrograms
	FindProdiByID(studyProgramID uint64) models.StudyPrograms
	FindProdiByCode(Code string) models.StudyPrograms
}

type studyProgramService struct {
	studyProgramRepository repository.StudyProgramRepository
}

func NewStudyProgramService(studyProgramRepo repository.StudyProgramRepository) StudyProgramService {
	return &studyProgramService{
		studyProgramRepository: studyProgramRepo,
	}
}

func (service *studyProgramService) InsertProdi(studyProgram dto.StudyProgramCreateDTO) models.StudyPrograms {
	studyProgramToInsert := models.StudyPrograms{}
	studyProgramToInsert.Code = studyProgram.Code
	studyProgramToInsert.Name = studyProgram.Name
	studyProgramInserted := service.studyProgramRepository.InsertStudyProgram(studyProgramToInsert)
	return studyProgramInserted
}

func (service *studyProgramService) UpdateProdi(studyProgram dto.StudyProgramUpdateDTO) models.StudyPrograms {
	studyProgramToUpdate := models.StudyPrograms{}
	studyProgramToUpdate.ID = uint64(studyProgram.ID)
	studyProgramToUpdate.Code = studyProgram.Code
	studyProgramToUpdate.Name = studyProgram.Name
	studyProgramUpdated := service.studyProgramRepository.UpdateStudyProgram(studyProgramToUpdate)
	return studyProgramUpdated
}

func (service *studyProgramService) DeleteProdi(id uint64) error {
	err := service.studyProgramRepository.DeleteStudyProgram(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *studyProgramService) AllProdi() []models.StudyPrograms {
	allStudyProgram := service.studyProgramRepository.AllStudyProgram()
	return allStudyProgram
}

func (service *studyProgramService) FindProdiByID(studyProgramID uint64) models.StudyPrograms {
	studyProgram := service.studyProgramRepository.FindStudyProgramByID(studyProgramID)
	return studyProgram
}

func (service *studyProgramService) FindProdiByCode(Code string) models.StudyPrograms {
	studyProgram := service.studyProgramRepository.FindStudyProgramByCode(Code)
	return studyProgram
}
