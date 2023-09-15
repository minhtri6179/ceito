// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: answer.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnswer = `-- name: CreateAnswer :one
INSERT INTO answer (question_id, answer_text, is_correct)
VALUES ($1, $2, $3)
RETURNING answer_id, question_id, answer_text, is_correct, created_at, update_at
`

type CreateAnswerParams struct {
	QuestionID pgtype.Int4 `json:"question_id"`
	AnswerText pgtype.Text `json:"answer_text"`
	IsCorrect  pgtype.Bool `json:"is_correct"`
}

func (q *Queries) CreateAnswer(ctx context.Context, arg CreateAnswerParams) (Answer, error) {
	row := q.db.QueryRow(ctx, createAnswer, arg.QuestionID, arg.AnswerText, arg.IsCorrect)
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

const deleteAnswer = `-- name: DeleteAnswer :exec
DELETE FROM answer
WHERE answer_id = $1
`

func (q *Queries) DeleteAnswer(ctx context.Context, answerID int64) error {
	_, err := q.db.Exec(ctx, deleteAnswer, answerID)
	return err
}

const getAnswer = `-- name: GetAnswer :one
SELECT answer_id, question_id, answer_text, is_correct, created_at, update_at
FROM answer
WHERE answer_id = $1
`

func (q *Queries) GetAnswer(ctx context.Context, answerID int64) (Answer, error) {
	row := q.db.QueryRow(ctx, getAnswer, answerID)
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

const listAnswers = `-- name: ListAnswers :many
SELECT answer_id, question_id, answer_text, is_correct, created_at, update_at
FROM answer
ORDER BY answer_id
LIMIT $1 OFFSET $2
`

type ListAnswersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAnswers(ctx context.Context, arg ListAnswersParams) ([]Answer, error) {
	rows, err := q.db.Query(ctx, listAnswers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Answer{}
	for rows.Next() {
		var i Answer
		if err := rows.Scan(
			&i.AnswerID,
			&i.QuestionID,
			&i.AnswerText,
			&i.IsCorrect,
			&i.CreatedAt,
			&i.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAnswer = `-- name: UpdateAnswer :exec
UPDATE answer
SET question_id = $1,
    answer_text = $2,
    is_correct = $3
WHERE answer_id = $4
`

type UpdateAnswerParams struct {
	QuestionID pgtype.Int4 `json:"question_id"`
	AnswerText pgtype.Text `json:"answer_text"`
	IsCorrect  pgtype.Bool `json:"is_correct"`
	AnswerID   int64       `json:"answer_id"`
}

func (q *Queries) UpdateAnswer(ctx context.Context, arg UpdateAnswerParams) error {
	_, err := q.db.Exec(ctx, updateAnswer,
		arg.QuestionID,
		arg.AnswerText,
		arg.IsCorrect,
		arg.AnswerID,
	)
	return err
}
