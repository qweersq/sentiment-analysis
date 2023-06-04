package dto

type SentimenAnalysisCreateDTO struct {
	SentimentType   string  `json:"sentiment_type" form:"sentiment_type" binding:"required"`
	CommentID       uint64  `json:"comment_id" form:"comment_id" binding:"required"`
	ConfidenceLevel float64 `json:"confidence_level" form:"confidence_level" binding:"required"`
}

type SentimenAnalysisUpdateDTO struct {
	ID              uint64  `json:"id" form:"id" binding:"required"`
	SentimentType   string  `json:"sentiment_type" form:"sentiment_type" binding:"required"`
	CommentID       uint64  `json:"comment_id" form:"comment_id" binding:"required"`
	ConfidenceLevel float64 `json:"confidence_level" form:"confidence_level" binding:"required"`
}

type SentimenAnalysisDTO struct {
	ID               uint64  `json:"id" form:"id" binding:"required"`
	SentimentType    string  `json:"sentiment_type" form:"sentiment_type" binding:"required"`
	ConfidenceLevel  float64 `json:"confidence_level" form:"confidence_level" binding:"required"`
	CommentID        uint64  `json:"comment_id" form:"comment_id" binding:"required"`
	Comment          string  `json:"comment" form:"comment" binding:"required"`
	CourseID         uint64  `json:"course_id" form:"course_id" binding:"required"`
	CourseCode       string  `json:"course_code" form:"course_code" binding:"required"`
	CourseName       string  `json:"course_name" form:"course_name" binding:"required"`
	LecturerID       uint64  `json:"lecturer_id" form:"lecturer_id" binding:"required"`
	LecturerCode     string  `json:"lecturer_code" form:"lecturer_code" binding:"required"`
	LecturerName     string  `json:"lecturer_name" form:"lecturer_name" binding:"required"`
	StudyProgramID   uint64  `json:"study_program_id" form:"study_program_id" binding:"required"`
	StudyProgramCode string  `json:"study_program_code" form:"study_program_code" binding:"required"`
	StudyProgramName string  `json:"study_program_name" form:"study_program_name" binding:"required"`
	SchoolYear       uint    `json:"school_year" form:"school_year" binding:"required"`
	Semester         uint    `json:"semester" form:"semester" binding:"required"`
}
