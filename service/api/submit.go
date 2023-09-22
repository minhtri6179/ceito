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
type submitResponse struct {
	Mark int `json:"your_score"`
}

func (server *Server) submitAnswer(ctx *gin.Context) {
	var req submitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	maring := 0
	for i, q := range req.QuestionID {
		res, err := server.checkQuestion(ctx, q, req.AnswerID[i])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		if res {
			maring++
		}
	}
	var res submitResponse
	var err error
	res.Mark, err = server.numofTrue2scoreListing(ctx, maring)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Print(maring)

	ctx.JSON(http.StatusOK, res)
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
func makeListeningScore(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i*5 + 10
	}
	a[0] = 0
	var s []int = a[97:]
	for i := 0; i < len(s); i++ {
		s[i] = 495
	}
	return a
}

func (server *Server) numofTrue2scoreListing(ctx *gin.Context, numofTrue int) (int, error) {
	mapping := makeListeningScore(0, 100)
	return mapping[numofTrue], nil
}
