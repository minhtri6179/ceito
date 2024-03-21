// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateAnswer(ctx context.Context, arg CreateAnswerParams) (Answer, error)
	CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error)
	CreateScore(ctx context.Context, arg CreateScoreParams) (Score, error)
	CreateTest(ctx context.Context, username pgtype.Text) (TestResult, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteAnswer(ctx context.Context, answerID int64) error
	DeleteQuestion(ctx context.Context, questionID int64) error
	DeleteScore(ctx context.Context, id int64) error
	DeleteTest(ctx context.Context, testID int64) error
	DeleteUser(ctx context.Context, username string) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetAnswer(ctx context.Context, answerID int64) (Answer, error)
	GetAnswerPart(ctx context.Context, questionID pgtype.Int4) ([]Answer, error)
	GetQuestion(ctx context.Context, questionID int64) (Question, error)
	GetScore(ctx context.Context, id int64) (Score, error)
	GetTest(ctx context.Context, testID int64) (TestResult, error)
	GetTestQuestions(ctx context.Context, testName pgtype.Text) ([]Question, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListAnswers(ctx context.Context, arg ListAnswersParams) ([]Answer, error)
	ListQuestions(ctx context.Context, arg ListQuestionsParams) ([]Question, error)
	ListScores(ctx context.Context, arg ListScoresParams) ([]Score, error)
	ListTests(ctx context.Context, arg ListTestsParams) ([]TestResult, error)
	UpdateAnswer(ctx context.Context, arg UpdateAnswerParams) error
	UpdateQuestion(ctx context.Context, arg UpdateQuestionParams) error
	UpdateScore(ctx context.Context, arg UpdateScoreParams) error
	UpdateTest(ctx context.Context, arg UpdateTestParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)