package service

import (
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"
)

type CommentService interface {
	InsertComment(comment dto.CommentCreateDTO) dto.CommentCreateDTO
	UpdateComment(comment dto.CommentUpdateDTO) dto.CommentUpdateDTO
	GetAllComment() []dto.CommentDTO
	GetCommentByID(id uint64) dto.CommentDTO
	DeleteComment(id uint64) error
}

type commentService struct {
	commentRepository  repository.CommentRepository
	lecturerRepository repository.LecturerRepository
}

func NewCommentServiceWithLecturer(commentRepo repository.CommentRepository, lecturerRepo repository.LecturerRepository) CommentService {
	return &commentService{
		commentRepository:  commentRepo,
		lecturerRepository: lecturerRepo,
	}
}

func (service *commentService) InsertComment(comment dto.CommentCreateDTO) dto.CommentCreateDTO {
	commentToInsert := models.Comments{}
	commentToInsert.Comment = comment.Comment
	commentToInsert.CourseID = uint64(comment.CourseID)
	commentToInsert.LecturerID = uint64(comment.LecturerID)
	commentToInsert.SchoolYear = comment.SchoolYear
	commentToInsert.Semester = comment.Semester
	commentInserted := service.commentRepository.InsertComment(commentToInsert)

	var commentDTO dto.CommentCreateDTO
	commentDTO.Comment = commentInserted.Comment
	commentDTO.CourseID = commentInserted.CourseID
	commentDTO.LecturerID = commentInserted.LecturerID
	commentDTO.SchoolYear = commentInserted.SchoolYear
	commentDTO.Semester = commentInserted.Semester
	return commentDTO
}

func (service *commentService) UpdateComment(comment dto.CommentUpdateDTO) dto.CommentUpdateDTO {
	commentToUpdate := models.Comments{}
	commentToUpdate.ID = uint64(comment.ID)
	commentToUpdate.Comment = comment.Comment
	commentToUpdate.CourseID = uint64(comment.CourseID)
	commentToUpdate.LecturerID = uint64(comment.LecturerID)
	commentToUpdate.SchoolYear = comment.SchoolYear
	commentToUpdate.Semester = comment.Semester
	commentUpdated := service.commentRepository.UpdateComment(commentToUpdate)

	var commentDTO dto.CommentUpdateDTO
	commentDTO.ID = commentUpdated.ID
	commentDTO.Comment = commentUpdated.Comment
	commentDTO.CourseID = commentUpdated.CourseID
	commentDTO.LecturerID = commentUpdated.LecturerID
	commentDTO.SchoolYear = commentUpdated.SchoolYear
	commentDTO.Semester = commentUpdated.Semester
	return commentDTO
}

func (service *commentService) GetAllComment() []dto.CommentDTO {
	allComment := service.commentRepository.GetAllComment()

	var commentDTO []dto.CommentDTO
	for _, value := range allComment {
		var comment dto.CommentDTO
		lecturerData := service.lecturerRepository.GetLecturerByID(value.LecturerID)
		comment.CommentID = value.ID
		comment.Comment = value.Comment
		comment.CourseID = value.CourseID
		comment.CourseCode = value.Course.Code
		comment.CourseName = value.Course.Name
		comment.LecturerID = value.LecturerID
		comment.LecturerCode = lecturerData.Code
		comment.LecturerName = lecturerData.Name
		comment.StudyProgramID = lecturerData.StudyProgramID
		comment.StudyProgramCode = lecturerData.StudyProgram.Code
		comment.StudyProgramName = lecturerData.StudyProgram.Name
		comment.SchoolYear = value.SchoolYear
		comment.Semester = value.Semester
		commentDTO = append(commentDTO, comment)
	}

	return commentDTO
}

func (service *commentService) GetCommentByID(id uint64) dto.CommentDTO {
	comment := service.commentRepository.GetCommentByID(id)

	var commentDTO dto.CommentDTO
	lecturerData := service.lecturerRepository.GetLecturerByID(comment.LecturerID)
	commentDTO.CommentID = comment.ID
	commentDTO.Comment = comment.Comment
	commentDTO.CourseID = comment.CourseID
	commentDTO.CourseCode = comment.Course.Code
	commentDTO.CourseName = comment.Course.Name
	commentDTO.LecturerID = comment.LecturerID
	commentDTO.LecturerCode = lecturerData.Code
	commentDTO.LecturerName = lecturerData.Name
	commentDTO.StudyProgramID = lecturerData.StudyProgramID
	commentDTO.StudyProgramCode = lecturerData.StudyProgram.Code
	commentDTO.StudyProgramName = lecturerData.StudyProgram.Name
	commentDTO.SchoolYear = comment.SchoolYear
	commentDTO.Semester = comment.Semester

	return commentDTO
}

func (service *commentService) DeleteComment(id uint64) error {
	err := service.commentRepository.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}
