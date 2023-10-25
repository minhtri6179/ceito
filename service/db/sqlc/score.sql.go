// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: score.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createScore = `-- name: CreateScore :one
INSERT INTO score (
        reading_score,
        listening_score,
        total_score
    )
VALUES ($1, $2, $3)
RETURNING id, reading_score, listening_score, total_score
`

type CreateScoreParams struct {
	ReadingScore   pgtype.Int4 `json:"reading_score"`
	ListeningScore pgtype.Int4 `json:"listening_score"`
	TotalScore     pgtype.Int4 `json:"total_score"`
}

func (q *Queries) CreateScore(ctx context.Context, arg CreateScoreParams) (Score, error) {
	row := q.db.QueryRow(ctx, createScore, arg.ReadingScore, arg.ListeningScore, arg.TotalScore)
	var i Score
	err := row.Scan(
		&i.ID,
		&i.ReadingScore,
		&i.ListeningScore,
		&i.TotalScore,
	)
	return i, err
}

const deleteScore = `-- name: DeleteScore :exec
DELETE FROM score
WHERE id = $1
`

func (q *Queries) DeleteScore(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteScore, id)
	return err
}

const getScore = `-- name: GetScore :one
SELECT id, reading_score, listening_score, total_score
FROM score
WHERE id = $1
`

func (q *Queries) GetScore(ctx context.Context, id int64) (Score, error) {
	row := q.db.QueryRow(ctx, getScore, id)
	var i Score
	err := row.Scan(
		&i.ID,
		&i.ReadingScore,
		&i.ListeningScore,
		&i.TotalScore,
	)
	return i, err
}

const listScores = `-- name: ListScores :many
SELECT id, reading_score, listening_score, total_score
FROM score
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListScoresParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListScores(ctx context.Context, arg ListScoresParams) ([]Score, error) {
	rows, err := q.db.Query(ctx, listScores, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Score{}
	for rows.Next() {
		var i Score
		if err := rows.Scan(
			&i.ID,
			&i.ReadingScore,
			&i.ListeningScore,
			&i.TotalScore,
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

const updateScore = `-- name: UpdateScore :exec
UPDATE score
SET reading_score = $1,
    listening_score = $2,
    total_score = $3
WHERE id = $4
`

type UpdateScoreParams struct {
	ReadingScore   pgtype.Int4 `json:"reading_score"`
	ListeningScore pgtype.Int4 `json:"listening_score"`
	TotalScore     pgtype.Int4 `json:"total_score"`
	ID             int64       `json:"id"`
}

func (q *Queries) UpdateScore(ctx context.Context, arg UpdateScoreParams) error {
	_, err := q.db.Exec(ctx, updateScore,
		arg.ReadingScore,
		arg.ListeningScore,
		arg.TotalScore,
		arg.ID,
	)
	return err
}
