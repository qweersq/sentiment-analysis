package dto

type StudyProgramDTO struct {
	ID   uint   `json:"id" form:"id" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}
type StudyProgramCreateDTO struct {
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

type StudyProgramUpdateDTO struct {
	ID   uint   `json:"id" form:"id" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

type StudyProgramIDDTO struct {
	ID uint64 `json:"id" form:"id"`
}

type StudyProgramCodeDTO struct {
	Code string `json:"code" form:"code" binding:"required"`
}

type ErrorDTO struct {
	Error ErrorStudyProgramDTO `json:"error"`
}

type ErrorStudyProgramDTO struct {
	Message string `json:"message"`
}
