package repository

import (
	"sentiment/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	InsertComment(comments models.Comments) models.Comments
	InsertToSentimentAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis
	UpdateComment(comments models.Comments) models.Comments
	GetAllComment() []models.Comments
	GetCommentByID(id uint64) models.Comments
	DeleteComment(id uint64) error
}

type commentConnection struct {
	connection *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentConnection{
		connection: db,
	}
}

func (db *commentConnection) InsertComment(comments models.Comments) models.Comments {
	db.connection.Save(&comments)
	return comments
}

func (db *commentConnection) UpdateComment(comments models.Comments) models.Comments {
	db.connection.Save(&comments)
	return comments
}

func (db *commentConnection) InsertToSentimentAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis {
	db.connection.Save(&sentimentAnalysis)
	return sentimentAnalysis
}

func (db *commentConnection) GetAllComment() []models.Comments {
	var comments []models.Comments
	db.connection.Preload("Course").Preload("Lecturer").Find(&comments)
	return comments
}

func (db *commentConnection) GetCommentByID(id uint64) models.Comments {
	var comment models.Comments
	db.connection.Preload("Course").Preload("Lecturer").Find(&comment, id)
	return comment
}

func (db *commentConnection) DeleteComment(id uint64) error {
	comment := models.Comments{}
	result := db.connection.Delete(&comment, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
