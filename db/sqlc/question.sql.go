// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: question.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createQuestion = `-- name: CreateQuestion :one
INSERT INTO question (
        question_text,
        test_name,
        img,
        created_at
    )
VALUES ($1, $2, $3, $4)
RETURNING question_id, question_text, test_name, img, created_at, update_at
`

type CreateQuestionParams struct {
	QuestionText pgtype.Text        `json:"question_text"`
	TestName     pgtype.Text        `json:"test_name"`
	Img          []byte             `json:"img"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error) {
	row := q.db.QueryRow(ctx, createQuestion,
		arg.QuestionText,
		arg.TestName,
		arg.Img,
		arg.CreatedAt,
	)
	var i Question
	err := row.Scan(
		&i.QuestionID,
		&i.QuestionText,
		&i.TestName,
		&i.Img,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteQuestion = `-- name: DeleteQuestion :exec
DELETE FROM question
WHERE question_id = $1
`

func (q *Queries) DeleteQuestion(ctx context.Context, questionID int64) error {
	_, err := q.db.Exec(ctx, deleteQuestion, questionID)
	return err
}

const getQuestion = `-- name: GetQuestion :one
SELECT question_id, question_text, test_name, img, created_at, update_at
FROM question
WHERE question_id = $1
`

func (q *Queries) GetQuestion(ctx context.Context, questionID int64) (Question, error) {
	row := q.db.QueryRow(ctx, getQuestion, questionID)
	var i Question
	err := row.Scan(
		&i.QuestionID,
		&i.QuestionText,
		&i.TestName,
		&i.Img,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const getTestQuestions = `-- name: GetTestQuestions :many
SELECT question_id, question_text, test_name, img, created_at, update_at
FROM question
WHERE test_name = $1
ORDER BY question_id
`

func (q *Queries) GetTestQuestions(ctx context.Context, testName pgtype.Text) ([]Question, error) {
	rows, err := q.db.Query(ctx, getTestQuestions, testName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Question{}
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.QuestionID,
			&i.QuestionText,
			&i.TestName,
			&i.Img,
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

const listQuestions = `-- name: ListQuestions :many
SELECT question_id, question_text, test_name, img, created_at, update_at
FROM question
ORDER BY question_id
LIMIT $1 OFFSET $2
`

type ListQuestionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListQuestions(ctx context.Context, arg ListQuestionsParams) ([]Question, error) {
	rows, err := q.db.Query(ctx, listQuestions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Question{}
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.QuestionID,
			&i.QuestionText,
			&i.TestName,
			&i.Img,
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

const updateQuestion = `-- name: UpdateQuestion :exec
UPDATE question
SET question_text = $1,
    test_name = $2
WHERE question_id = $3
`

type UpdateQuestionParams struct {
	QuestionText pgtype.Text `json:"question_text"`
	TestName     pgtype.Text `json:"test_name"`
	QuestionID   int64       `json:"question_id"`
}

func (q *Queries) UpdateQuestion(ctx context.Context, arg UpdateQuestionParams) error {
	_, err := q.db.Exec(ctx, updateQuestion, arg.QuestionText, arg.TestName, arg.QuestionID)
	return err
}