package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/minhtri6179/service/db/sqlc"
)

type createAnswerRequest struct {
	QuestionID pgtype.Int4 `json:"question_id" binding:"required,number"`
	AnswerText pgtype.Text `json:"answer_text" binding:"required"`
	IsCorrect  pgtype.Bool `json:"is_correct"`
}

func (server *Server) createAnswer(ctx *gin.Context) {
	var req createAnswerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateAnswerParams{
		QuestionID: req.QuestionID,
		AnswerText: req.AnswerText,
		IsCorrect:  req.IsCorrect,
	}
	user, err := server.store.CreateAnswer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

type updateAnswerRequest struct {
	QuestionID pgtype.Int4 `json:"question_id"`
	AnswerText pgtype.Text `json:"answer_text"`
	IsCorrect  pgtype.Bool `json:"is_correct"`
	AnswerID   int64       `uri:"id"`
}

func (server *Server) updateAnswer(ctx *gin.Context) {
	var req updateAnswerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.UpdateAnswerParams{
		QuestionID: req.QuestionID,
		AnswerText: req.AnswerText,
		IsCorrect:  req.IsCorrect,
		AnswerID:   req.AnswerID,
	}

	err := server.store.UpdateAnswer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Update question successfully")

}

type listAnswerRequest struct {
	TestName pgtype.Text `json:"test_name"`
}

func (server *Server) listAnswers(ctx *gin.Context) {
	var req listAnswerRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAnswersParams{
		Limit:  400,
		Offset: 0,
	}
	answers, err := server.store.ListAnswers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, answers)
}
