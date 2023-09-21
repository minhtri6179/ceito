package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type submitRequest struct {
	QuestionID []int64 `json:"question_id" `
	AnswerID   []int64 `json:"answer_id" `
}

func (server *Server) submitAnswer(ctx *gin.Context) {
	var req submitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// if len(req.QuestionID) != len(req.AnswerID) {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }

	// for i := 0; i < len(req.QuestionID); i++ {
	// 	server.checkQuestion(ctx, req.QuestionID[i], req.AnswerID[i])
	// }
	res, err := server.checkQuestion(ctx, req.QuestionID[0], req.AnswerID[0])
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Print(res)

	ctx.JSON(http.StatusOK, true)
}
func (server *Server) checkQuestion(ctx *gin.Context, questionID int64, answerID int64) (bool, error) {
	answer, err := server.store.GetAnswer(ctx, answerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false, err
	}
	if int(answer.QuestionID.Int32) != int(questionID) {
		err := fmt.Errorf("answer %d does not belong to question %d", answerID, questionID)
		return false, err
	}

	if answer.IsCorrect.Bool {
		return true, nil
	}

	return false, nil
}
