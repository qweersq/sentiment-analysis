package service

import (
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"
)

type CourseService interface {
	InsertCourse(course dto.CourseCreateDTO) dto.CourseCreateDTO
	UpdateCourse(course dto.CourseUpdateDTO) dto.CourseUpdateDTO
	GetAllCourse() []models.Courses
	GetCourseByID(id uint64) models.Courses
	DeleteCourse(id uint64) error
}

type courseService struct {
	courseRepository repository.CourseRepository
}

func NewCourseService(courseRepo repository.CourseRepository) CourseService {
	return &courseService{
		courseRepository: courseRepo,
	}
}

func (service *courseService) InsertCourse(course dto.CourseCreateDTO) dto.CourseCreateDTO {
	courseToInsert := models.Courses{}
	courseToInsert.Code = course.Code
	courseToInsert.Name = course.Name
	courseInserted := service.courseRepository.InsertCourse(courseToInsert)

	var courseDTO dto.CourseCreateDTO
	courseDTO.Code = courseInserted.Code
	courseDTO.Name = courseInserted.Name
	return courseDTO
}

func (service *courseService) UpdateCourse(course dto.CourseUpdateDTO) dto.CourseUpdateDTO {
	courseToUpdate := models.Courses{}
	courseToUpdate.ID = uint64(course.ID)
	courseToUpdate.Code = course.Code
	courseToUpdate.Name = course.Name
	courseUpdated := service.courseRepository.UpdateCourse(courseToUpdate)

	var courseDTO dto.CourseUpdateDTO
	courseDTO.ID = courseUpdated.ID
	courseDTO.Code = courseUpdated.Code
	courseDTO.Name = courseUpdated.Name
	return courseDTO
}

func (service *courseService) GetAllCourse() []models.Courses {
	allCourse := service.courseRepository.GetAllCourse()
	return allCourse
}

func (service *courseService) GetCourseByID(id uint64) models.Courses {
	course := service.courseRepository.GetCourseByID(id)
	return course
}

func (service *courseService) DeleteCourse(id uint64) error {
	err := service.courseRepository.DeleteCourse(id)
	if err != nil {
		return err
	}
	return nil
}
