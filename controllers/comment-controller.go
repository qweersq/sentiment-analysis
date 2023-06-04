package controllers

import (
	"net/http"
	"sentiment/dto"
	"sentiment/helper"
	"sentiment/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAllData(ctx *gin.Context)
	GetDataByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type commentController struct {
	commentService service.CommentService
	jwtService     service.JWTService
}

func NewCommentController(commentService service.CommentService, jwtService service.JWTService) CommentController {
	return &commentController{
		commentService: commentService,
		jwtService:     jwtService,
	}
}

func (c *commentController) Create(ctx *gin.Context) {
	var commentCreate dto.CommentCreateDTO
	errDTO := ctx.ShouldBind(&commentCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	comment := c.commentService.InsertComment(commentCreate)
	res := helper.BuildResponse(true, "OK!", comment)
	ctx.JSON(http.StatusOK, res)
}

func (c *commentController) Update(ctx *gin.Context) {
	var commentUpdateDTO dto.CommentUpdateDTO
	errDTO := ctx.ShouldBind(&commentUpdateDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	comment := c.commentService.UpdateComment(commentUpdateDTO)
	res := helper.BuildResponse(true, "OK!", comment)
	ctx.JSON(http.StatusOK, res)
}

func (c *commentController) GetAllData(ctx *gin.Context) {
	comments := c.commentService.GetAllComment()
	res := helper.BuildResponse(true, "OK!", comments)
	ctx.JSON(http.StatusOK, res)
}

func (c *commentController) GetDataByID(ctx *gin.Context) {
	id := ctx.Param("id")
	commentID, _ := strconv.ParseUint(id, 0, 0)
	comment := c.commentService.GetCommentByID(commentID)
	if comment.CommentID == 0 {
		res := helper.BuildErrorResponse("Data not found", "Data not found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", comment)
	ctx.JSON(http.StatusOK, res)
}

func (c *commentController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	commentID, _ := strconv.ParseUint(id, 0, 0)
	err := c.commentService.DeleteComment(commentID)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
