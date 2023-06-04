package controllers

import (
	"net/http"
	"sentiment/dto"
	"sentiment/helper"
	"sentiment/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudyProgramController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindByCode(ctx *gin.Context)
}

type studyProgramController struct {
	studyProgramService service.StudyProgramService
	jwtService          service.JWTService
}

func NewStudyProgramController(studyProgramService service.StudyProgramService, jwtService service.JWTService) StudyProgramController {
	return &studyProgramController{
		studyProgramService: studyProgramService,
		jwtService:          jwtService,
	}
}

func (c *studyProgramController) Create(ctx *gin.Context) {
	// create study program
	var studyProgramCreateDTO dto.StudyProgramCreateDTO
	errDTO := ctx.ShouldBind(&studyProgramCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	studyProgram := c.studyProgramService.InsertProdi(studyProgramCreateDTO)
	res := helper.BuildResponse(true, "OK!", studyProgram)
	ctx.JSON(http.StatusOK, res)
}

func (c *studyProgramController) Update(ctx *gin.Context) {
	var studyProgramUpdateDTO dto.StudyProgramUpdateDTO
	errDTO := ctx.ShouldBind(&studyProgramUpdateDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	studyProgram := c.studyProgramService.UpdateProdi(studyProgramUpdateDTO)
	res := helper.BuildResponse(true, "OK!", studyProgram)
	ctx.JSON(http.StatusOK, res)
}

func (c *studyProgramController) Delete(ctx *gin.Context) {
	studyProgramID := ctx.Param("id")

	// Konversi studyProgramID menjadi tipe data uint64 atau sesuai tipe ID yang digunakan
	id, err := strconv.ParseUint(studyProgramID, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid study program ID", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = c.studyProgramService.DeleteProdi(id)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete study program", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "Study program deleted successfully", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (c *studyProgramController) All(ctx *gin.Context) {
	studyPrograms := c.studyProgramService.AllProdi()
	res := helper.BuildResponse(true, "OK!", studyPrograms)
	ctx.JSON(http.StatusOK, res)
}

func (c *studyProgramController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	// Konversi studyProgramID menjadi tipe data uint64 atau sesuai tipe ID yang digunakan
	studyProgramID, _ := strconv.ParseUint(id, 10, 64)
	studyProgram, err := c.studyProgramService.FindProdiByID(studyProgramID)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get study program", "No data found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK!", studyProgram)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *studyProgramController) FindByCode(ctx *gin.Context) {
	var studyProgramCode dto.StudyProgramCodeDTO
	errDTO := ctx.ShouldBind(&studyProgramCode)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	studyPrograms := c.studyProgramService.FindProdiByCode(studyProgramCode.Code)
	res := helper.BuildResponse(true, "OK!", studyPrograms)
	ctx.JSON(http.StatusOK, res)
}
