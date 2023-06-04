package dto

type LectureCreateDTO struct {
	Code           string `json:"code" form:"code" binding:"required"`
	Name           string `json:"name" form:"name" binding:"required"`
	StudyProgramID uint64 `json:"study_program_id" form:"study_program_id" binding:"required"`
}
