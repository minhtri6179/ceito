package api

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	db "github.com/minhtri6179/service/db/sqlc"
	"github.com/minhtri6179/service/token"
	"github.com/minhtri6179/service/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		MaxAge:       12 * time.Hour,
	}))

	router.Use(cors.Default())

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	// Question
	router.POST("/questions", server.createQuestion)
	router.PUT("/questions/:id", server.updateQuestion)
	router.GET("/questions/:name", server.listQuestions)
	// Answer
	router.POST("/answers", server.createAnswer)
	router.PUT("/answers/:id", server.updateAnswer)
	router.GET("/answers", server.listAnswers)
	router.GET("/answers-part/:name", server.getAnswerPart)
	// Score
	router.POST("/scores", server.createScore)
	router.PUT("/scores/:id", server.updateScore)
	// Test
	router.POST("/tests", server.createTest)
	router.PUT("/tests/:id", server.updateTest)
	// Submit
	router.POST("/submit", server.submitAnswer)
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
