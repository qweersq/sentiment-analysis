package dto

type CourseCreateDTO struct {
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

type CourseUpdateDTO struct {
	ID   uint64 `json:"id" form:"id" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}
