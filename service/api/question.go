package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/minhtri6179/service/db/sqlc"
)

type createQuestionRequest struct {
	QuestionText pgtype.Text       `json:"question_text"`
	AnswerID     pgtype.Int4       `json:"answer_id"`
	TestName     pgtype.Text       `json:"test_name"`
	Image        pgtype.JSONBCodec `json:"image"`
}

func (server *Server) createQuestion(ctx *gin.Context) {
	var req createQuestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jsonData, err := json.Marshal(req.Image)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateQuestionParams{
		QuestionText: req.QuestionText,
		TestName:     req.TestName,
		Img:          jsonData,
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
	TestName string `uri:"name"`
}

func (server *Server) listQuestions(ctx *gin.Context) {
	var req listQuestionsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	textValue := pgtype.Text{}
	textValue.String = req.TestName
	textValue.Valid = true

	questions, err := server.store.GetTestQuestions(ctx, textValue)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, questions)
}
