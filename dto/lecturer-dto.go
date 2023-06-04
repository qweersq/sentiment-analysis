package dto

type LecturerCreateDTO struct {
	Code           string `json:"code" form:"code" binding:"required"`
	Name           string `json:"name" form:"name" binding:"required"`
	StudyProgramID uint64 `json:"study_program_id" form:"study_program_id" binding:"required"`
}

type LecturerUpdateDTO struct {
	ID             uint64 `json:"id" form:"id" binding:"required"`
	Code           string `json:"code" form:"code" binding:"required"`
	Name           string `json:"name" form:"name" binding:"required"`
	StudyProgramID uint64 `json:"study_program_id" form:"study_program_id" binding:"required"`
}
