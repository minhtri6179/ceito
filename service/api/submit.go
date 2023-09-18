package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	QuestionID int64 `json:"question_id" binding:"required,min=1"`
	AnswerID   int64 `json:"answer_id" binding:"required,min=1"`
}

func (server *Server) submitAnswer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := server.checkQuestion(ctx, req.QuestionID, req.AnswerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
func (server *Server) checkQuestion(ctx *gin.Context, questionID int64, answerID int64) (bool, error) {
	question, err := server.store.GetQuestion(ctx, questionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false, err
	}
	isTrue, err := server.store.GetAnswer(ctx, answerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false, err
	}

	if isTrue.IsCorrect.Bool && int64(question.AnswerID.Int32) == answerID {
		fmt.Printf("Correct question")
		return true, nil
	}
	return false, nil
}
