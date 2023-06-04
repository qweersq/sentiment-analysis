package service

import (
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"
)

type LecturerService interface {
	InsertLecturer(lecturer dto.LecturerCreateDTO) dto.LecturerCreateDTO
	UpdateLecturer(lecturer dto.LecturerUpdateDTO) dto.LecturerUpdateDTO
	GetAllLecturer() []models.Lecturers
	GetLecturerByID(id uint64) models.Lecturers
	DeleteLecturer(id uint64) error
}

type lecturerService struct {
	lecturerRepository repository.LecturerRepository
}

func NewLecturerService(lecturerRepo repository.LecturerRepository) LecturerService {
	return &lecturerService{
		lecturerRepository: lecturerRepo,
	}
}

func (service *lecturerService) InsertLecturer(lecturer dto.LecturerCreateDTO) dto.LecturerCreateDTO {
	lecturerToInsert := models.Lecturers{}
	lecturerToInsert.Code = lecturer.Code
	lecturerToInsert.Name = lecturer.Name
	lecturerToInsert.StudyProgramID = lecturer.StudyProgramID
	lecturerInserted := service.lecturerRepository.InsertLecturer(lecturerToInsert)

	var lecturerDTO dto.LecturerCreateDTO
	lecturerDTO.Code = lecturerInserted.Code
	lecturerDTO.Name = lecturerInserted.Name
	lecturerDTO.StudyProgramID = lecturerInserted.StudyProgramID
	return lecturerDTO
}

func (service *lecturerService) UpdateLecturer(lecturer dto.LecturerUpdateDTO) dto.LecturerUpdateDTO {
	lecturerToUpdate := models.Lecturers{}
	lecturerToUpdate.ID = uint64(lecturer.ID)
	lecturerToUpdate.Code = lecturer.Code
	lecturerToUpdate.Name = lecturer.Name
	lecturerToUpdate.StudyProgramID = lecturer.StudyProgramID
	lecturerUpdated := service.lecturerRepository.UpdateLecturer(lecturerToUpdate)

	var lecturerDTO dto.LecturerUpdateDTO
	lecturerDTO.ID = lecturerUpdated.ID
	lecturerDTO.Code = lecturerUpdated.Code
	lecturerDTO.Name = lecturerUpdated.Name
	lecturerDTO.StudyProgramID = lecturerUpdated.StudyProgramID
	return lecturerDTO
}

func (service *lecturerService) GetAllLecturer() []models.Lecturers {
	allLecturer := service.lecturerRepository.GetAllLecturer()
	return allLecturer
}

func (service *lecturerService) GetLecturerByID(id uint64) models.Lecturers {
	lecturer := service.lecturerRepository.GetLecturerByID(id)
	return lecturer
}

func (service *lecturerService) DeleteLecturer(id uint64) error {
	err := service.lecturerRepository.DeleteLecturer(id)
	if err != nil {
		return err
	}
	return nil
}
