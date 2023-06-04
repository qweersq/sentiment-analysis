package controllers

import (
	"net/http"
	"sentiment/dto"
	"sentiment/helper"
	"sentiment/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SentimenAnalysisController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAllData(ctx *gin.Context)
	GetDataByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type sentimenAnalysisController struct {
	sentimenAnalysisService service.SentimenAnalysisService
	jwtService              service.JWTService
}

func NewSentimenAnalysisController(sentimenAnalysisService service.SentimenAnalysisService, jwtService service.JWTService) SentimenAnalysisController {
	return &sentimenAnalysisController{
		sentimenAnalysisService: sentimenAnalysisService,
		jwtService:              jwtService,
	}
}

func (c *sentimenAnalysisController) Create(ctx *gin.Context) {
	var sentimenAnalysisCreate dto.SentimenAnalysisCreateDTO
	errDTO := ctx.ShouldBind(&sentimenAnalysisCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	sentimenAnalysis := c.sentimenAnalysisService.InsertSentimenAnalysis(sentimenAnalysisCreate)
	res := helper.BuildResponse(true, "OK!", sentimenAnalysis)
	ctx.JSON(http.StatusOK, res)
}

func (c *sentimenAnalysisController) Update(ctx *gin.Context) {
	var sentimenAnalysisUpdateDTO dto.SentimenAnalysisUpdateDTO
	errDTO := ctx.ShouldBind(&sentimenAnalysisUpdateDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	sentimenAnalysis := c.sentimenAnalysisService.UpdateSentimenAnalysis(sentimenAnalysisUpdateDTO)
	res := helper.BuildResponse(true, "OK!", sentimenAnalysis)
	ctx.JSON(http.StatusOK, res)
}

func (c *sentimenAnalysisController) GetAllData(ctx *gin.Context) {
	sentimenAnalysiss := c.sentimenAnalysisService.GetAllSentimenAnalysis()
	res := helper.BuildResponse(true, "OK!", sentimenAnalysiss)
	ctx.JSON(http.StatusOK, res)
}

func (c *sentimenAnalysisController) GetDataByID(ctx *gin.Context) {
	id := ctx.Param("id")
	sentimenAnalysisID, _ := strconv.ParseUint(id, 0, 0)
	sentimenAnalysis := c.sentimenAnalysisService.GetSentimenAnalysisByID(sentimenAnalysisID)
	if sentimenAnalysis.ID == 0 {
		res := helper.BuildErrorResponse("Data not found", "Data not found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", sentimenAnalysis)
	ctx.JSON(http.StatusOK, res)
}

func (c *sentimenAnalysisController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	sentimenAnalysisID, _ := strconv.ParseUint(id, 0, 0)
	err := c.sentimenAnalysisService.DeleteSentimenAnalysis(sentimenAnalysisID)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
