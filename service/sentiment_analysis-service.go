package service

import (
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"
)

type SentimenAnalysisService interface {
	InsertSentimenAnalysis(sentimenAnalysis dto.SentimenAnalysisCreateDTO) dto.SentimenAnalysisCreateDTO
	UpdateSentimenAnalysis(sentimenAnalysis dto.SentimenAnalysisUpdateDTO) dto.SentimenAnalysisUpdateDTO
	GetAllSentimenAnalysis() []models.SentimentAnalysis
	GetSentimenAnalysisByID(id uint64) models.SentimentAnalysis
	DeleteSentimenAnalysis(id uint64) error
}

type sentimenAnalysisService struct {
	sentimenAnalysisRepository repository.SentimenAnalysisRepository
	commentRepository          repository.CommentRepository
}

func NewSentimenAnalysisServiceWithComment(sentimenAnalysisRepo repository.SentimenAnalysisRepository, commentRepo repository.CommentRepository) SentimenAnalysisService {
	return &sentimenAnalysisService{
		sentimenAnalysisRepository: sentimenAnalysisRepo,
		commentRepository:          commentRepo,
	}
}

func (service *sentimenAnalysisService) InsertSentimenAnalysis(sentimenAnalysis dto.SentimenAnalysisCreateDTO) dto.SentimenAnalysisCreateDTO {
	sentimenAnalysisToInsert := models.SentimentAnalysis{}
	sentimenAnalysisToInsert.SentimentType = sentimenAnalysis.SentimentType
	sentimenAnalysisToInsert.CommentID = uint(sentimenAnalysis.CommentID)
	sentimenAnalysisToInsert.ConfidenceLevel = sentimenAnalysis.ConfidenceLevel
	sentimenAnalysisInserted := service.sentimenAnalysisRepository.InsertSentimenAnalysis(sentimenAnalysisToInsert)

	var sentimenAnalysisDTO dto.SentimenAnalysisCreateDTO
	sentimenAnalysisDTO.SentimentType = sentimenAnalysisInserted.SentimentType
	sentimenAnalysisDTO.ConfidenceLevel = sentimenAnalysisInserted.ConfidenceLevel
	sentimenAnalysisDTO.CommentID = uint64(sentimenAnalysisInserted.CommentID)

	return sentimenAnalysisDTO
}

func (service *sentimenAnalysisService) UpdateSentimenAnalysis(sentimenAnalysis dto.SentimenAnalysisUpdateDTO) dto.SentimenAnalysisUpdateDTO {
	sentimenAnalysisToUpdate := models.SentimentAnalysis{}
	sentimenAnalysisToUpdate.ID = uint64(sentimenAnalysis.ID)
	sentimenAnalysisToUpdate.SentimentType = sentimenAnalysis.SentimentType
	sentimenAnalysisToUpdate.CommentID = uint(sentimenAnalysis.CommentID)
	sentimenAnalysisToUpdate.ConfidenceLevel = sentimenAnalysis.ConfidenceLevel
	sentimenAnalysisUpdated := service.sentimenAnalysisRepository.UpdateSentimenAnalysis(sentimenAnalysisToUpdate)

	var sentimenAnalysisDTO dto.SentimenAnalysisUpdateDTO
	sentimenAnalysisDTO.ID = sentimenAnalysisUpdated.ID
	sentimenAnalysisDTO.SentimentType = sentimenAnalysisUpdated.SentimentType
	sentimenAnalysisDTO.CommentID = uint64(sentimenAnalysisUpdated.CommentID)
	sentimenAnalysisDTO.ConfidenceLevel = sentimenAnalysisUpdated.ConfidenceLevel
	return sentimenAnalysisDTO
}

func (service *sentimenAnalysisService) GetAllSentimenAnalysis() []models.SentimentAnalysis {
	allSentimenAnalysis := service.sentimenAnalysisRepository.GetAllSentimenAnalysis()
	return allSentimenAnalysis
}

func (service *sentimenAnalysisService) GetSentimenAnalysisByID(id uint64) models.SentimentAnalysis {
	sentimenAnalysis := service.sentimenAnalysisRepository.GetSentimenAnalysisByID(id)
	return sentimenAnalysis
}

func (service *sentimenAnalysisService) DeleteSentimenAnalysis(id uint64) error {
	err := service.sentimenAnalysisRepository.DeleteSentimenAnalysis(id)
	if err != nil {
		return err
	}
	return nil
}
