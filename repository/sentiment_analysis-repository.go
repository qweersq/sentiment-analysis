package repository

import (
	"sentiment/models"

	"gorm.io/gorm"
)

type SentimenAnalysisRepository interface {
	InsertSentimenAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis
	UpdateSentimenAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis
	GetAllSentimenAnalysis() []models.SentimentAnalysis
	GetSentimenAnalysisByID(id uint64) models.SentimentAnalysis
	DeleteSentimenAnalysis(id uint64) error
}

type sentimenAnalysisConnection struct {
	connection *gorm.DB
}

func NewSentimenAnalysisRepository(db *gorm.DB) SentimenAnalysisRepository {
	return &sentimenAnalysisConnection{
		connection: db,
	}
}

func (db *sentimenAnalysisConnection) InsertSentimenAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis {
	db.connection.Save(&sentimentAnalysis)
	return sentimentAnalysis
}

func (db *sentimenAnalysisConnection) UpdateSentimenAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis {
	db.connection.Save(&sentimentAnalysis)
	return sentimentAnalysis
}

func (db *sentimenAnalysisConnection) GetAllSentimenAnalysis() []models.SentimentAnalysis {
	var sentimentAnalysis []models.SentimentAnalysis
	db.connection.Preload("Comment").Preload("Comment.Course").Preload("Comment.Lecturer").Preload("Comment.Lecturer.StudyProgram").Find(&sentimentAnalysis)
	return sentimentAnalysis
}

func (db *sentimenAnalysisConnection) GetSentimenAnalysisByID(id uint64) models.SentimentAnalysis {
	var sentimenAnalysis models.SentimentAnalysis
	db.connection.Preload("Comment").Preload("Comment.Course").Preload("Comment.Lecturer").Preload("Comment.Lecturer.StudyProgram").Find(&sentimenAnalysis, id)
	return sentimenAnalysis
}

func (db *sentimenAnalysisConnection) DeleteSentimenAnalysis(id uint64) error {
	sentimenAnalysis := models.SentimentAnalysis{}
	result := db.connection.Delete(&sentimenAnalysis, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
