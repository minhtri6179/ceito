// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: answer.sql

package db

import (
	"context"
	"database/sql"
)

const createAnswer = `-- name: createAnswer :one
INSERT INTO answer (question_id, answer_text, is_correct)
VALUES ($1, $2, $3)
RETURNING answer_id, question_id, answer_text, is_correct, created_at, update_at
`

type createAnswerParams struct {
	QuestionID sql.NullInt32
	AnswerText sql.NullString
	IsCorrect  sql.NullBool
}

func (q *Queries) createAnswer(ctx context.Context, arg createAnswerParams) (Answer, error) {
	row := q.db.QueryRowContext(ctx, createAnswer, arg.QuestionID, arg.AnswerText, arg.IsCorrect)
	var i Answer
	err := row.Scan(
		&i.AnswerID,
		&i.QuestionID,
		&i.AnswerText,
		&i.IsCorrect,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}
