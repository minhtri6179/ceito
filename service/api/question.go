package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/minhtri6179/service/db/sqlc"
)

type createQuestionRequest struct {
	QuestionText pgtype.Text `json:"question_text"`
	AnswerID     pgtype.Int4 `json:"answer_id"`
	TestName     pgtype.Text `json:"test_name"`
}

func (server *Server) createQuestion(ctx *gin.Context) {
	var req createQuestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateQuestionParams{
		QuestionText: req.QuestionText,
		TestName:     req.TestName,
	}
	user, err := server.store.CreateQuestion(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

type updateQuestionRequest struct {
	QuestionText pgtype.Text `json:"question_text"`
	AnswerID     pgtype.Int4 `json:"answer_id" binding:"required,number"`
	ID           int64       `uri:"id"`
}

func (server *Server) updateQuestion(ctx *gin.Context) {
	var req updateQuestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	te, err := server.store.GetQuestion(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	arg := db.UpdateQuestionParams{
		QuestionText: te.QuestionText,
		QuestionID:   req.ID,
	}
	fmt.Printf("arg: %v", arg)

	err = server.store.UpdateQuestion(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Update question successfully")

}

type listQuestionsRequest struct {
	TestName pgtype.Text `json:"test_name"`
}

func (server *Server) listQuestions(ctx *gin.Context) {
	var req listQuestionsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListQuestionsParams{
		Limit:  100,
		Offset: 0,
	}
	questions, err := server.store.ListQuestions(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, questions)
}
