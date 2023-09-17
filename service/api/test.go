package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/minhtri6179/service/db/sqlc"
)

type createTestRequest struct {
	Username pgtype.Text `json:"username" validate:"required"`
}

func (server *Server) createTest(ctx *gin.Context) {
	var req createTestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := server.store.CreateTest(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

type updateTestRequest struct {
	Username pgtype.Text `json:"username"`
	TestID   int64       `uri:"id"`
}

func (server *Server) updateTest(ctx *gin.Context) {
	var req updateTestRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.UpdateTestParams{
		Username: req.Username,
		TestID:   req.TestID,
	}
	fmt.Printf("arg: %v", arg)

	err := server.store.UpdateTest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Update score successfully")

}
