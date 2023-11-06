package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/minhtri6179/service/db/sqlc"
)

type submitRequest struct {
	QuestionID []int64 `json:"question_id" `
	AnswerID   []int64 `json:"answer_id" `
}
type submitResponse struct {
	Mark int `json:"your_score"`
}
type updateScoreReq struct {
	ReadingScore   pgtype.Int4 `json:"reading_score" `
	ListeningScore pgtype.Int4 `json:"listening_score" `
	TotalScore     pgtype.Int4 `json:"total_score" `
}

func (server *Server) submitAnswer(ctx *gin.Context) {
	var req submitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var listen_score, read_score int = 0, 0

	for i, q := range req.QuestionID {
		res, err := server.checkQuestion(ctx, q, req.AnswerID[i])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		if res {
			if i < 100 {
				listen_score += 1
			} else {
				read_score += 1
			}
		}

	}
	var res submitResponse
	var err error
	listen_score, err = server.numofTrue2scoreListing(ctx, listen_score)
	read_score = server.numofTrue2scoreReading(ctx, read_score)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// Update score to database
	var socrereq updateScoreReq
	socrereq.ReadingScore.Valid = true
	socrereq.ReadingScore.Int32 = int32(read_score)
	socrereq.ListeningScore.Valid = true
	socrereq.ListeningScore.Int32 = int32(listen_score)
	socrereq.TotalScore.Valid = true
	socrereq.TotalScore.Int32 = int32(read_score + listen_score)
	res.Mark = int(socrereq.TotalScore.Int32)
	arg := db.CreateScoreParams{
		ReadingScore:   socrereq.ReadingScore,
		ListeningScore: socrereq.ListeningScore,
		TotalScore:     socrereq.TotalScore,
	}
	score, err := server.store.CreateScore(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println(score)

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) checkQuestion(ctx *gin.Context, questionID int64, answerID int64) (bool, error) {
	if answerID == 0 {
		return false, nil
	}
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
func (server *Server) numofTrue2scoreReading(ctx *gin.Context, numofTrue int) int {
	var res int = 0
	if numofTrue == 0 {
		res = 0

	} else if numofTrue == 1 {
		res = 5
	} else {
		res = (numofTrue - 1) * 5
	}
	return res

}

func (server *Server) numofTrue2scoreListing(ctx *gin.Context, numofTrue int) (int, error) {
	mapping := makeListeningScore(0, 100)
	return mapping[numofTrue], nil
}
