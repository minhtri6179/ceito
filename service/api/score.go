package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/minhtri6179/service/db/sqlc"
)

type createScoreRequest struct {
	TestID         pgtype.Int4 `json:"test_id" validate:"required"`
	ReadingScore   pgtype.Int4 `json:"reading_score"`
	ListeningScore pgtype.Int4 `json:"listening_score"`
	TotalScore     pgtype.Int4 `json:"total_score"`
}

func (server *Server) createScore(ctx *gin.Context) {
	var req createScoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateScoreParams{
		TestID:         req.TestID,
		ReadingScore:   req.ReadingScore,
		ListeningScore: req.ListeningScore,
		TotalScore:     req.TotalScore,
	}
	user, err := server.store.CreateScore(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

type updateScoreRequest struct {
	ScoreID        int64       `uri:"id"`
	TestID         pgtype.Int4 `json:"test_id" validate:"required"`
	ReadingScore   pgtype.Int4 `json:"reading_score"`
	TotalScore     pgtype.Int4 `json:"total_score"`
	ListeningScore pgtype.Int4 `json:"listening_score"`
}

func (server *Server) updateScore(ctx *gin.Context) {
	var req updateScoreRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.UpdateScoreParams{
		ScoreID:        req.ScoreID,
		TestID:         req.TestID,
		ReadingScore:   req.ReadingScore,
		ListeningScore: req.ListeningScore,
		TotalScore:     req.TotalScore,
	}
	fmt.Printf("arg: %v", arg)

	err := server.store.UpdateScore(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Update score successfully")

}
