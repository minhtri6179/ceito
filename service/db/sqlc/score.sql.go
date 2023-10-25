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
        test_id,
        reading_score,
        listening_score,
        total_score
    )
VALUES ($1, $2, $3, $4)
RETURNING score_id, test_id, reading_score, listening_score, total_score
`

type CreateScoreParams struct {
	TestID         pgtype.Int4 `json:"test_id"`
	ReadingScore   pgtype.Int4 `json:"reading_score"`
	ListeningScore pgtype.Int4 `json:"listening_score"`
	TotalScore     pgtype.Int4 `json:"total_score"`
}

func (q *Queries) CreateScore(ctx context.Context, arg CreateScoreParams) (Score, error) {
	row := q.db.QueryRow(ctx, createScore,
		arg.TestID,
		arg.ReadingScore,
		arg.ListeningScore,
		arg.TotalScore,
	)
	var i Score
	err := row.Scan(
		&i.ScoreID,
		&i.TestID,
		&i.ReadingScore,
		&i.ListeningScore,
		&i.TotalScore,
	)
	return i, err
}

const deleteScore = `-- name: DeleteScore :exec
DELETE FROM score
WHERE score_id = $1
`

func (q *Queries) DeleteScore(ctx context.Context, scoreID int64) error {
	_, err := q.db.Exec(ctx, deleteScore, scoreID)
	return err
}

const getScore = `-- name: GetScore :one
SELECT score_id, test_id, reading_score, listening_score, total_score
FROM score
WHERE score_id = $1
`

func (q *Queries) GetScore(ctx context.Context, scoreID int64) (Score, error) {
	row := q.db.QueryRow(ctx, getScore, scoreID)
	var i Score
	err := row.Scan(
		&i.ScoreID,
		&i.TestID,
		&i.ReadingScore,
		&i.ListeningScore,
		&i.TotalScore,
	)
	return i, err
}

const listScores = `-- name: ListScores :many
SELECT score_id, test_id, reading_score, listening_score, total_score
FROM score
ORDER BY score_id
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
			&i.ScoreID,
			&i.TestID,
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
SET test_id = $1,
    reading_score = $2,
    listening_score = $3,
    total_score = $4
WHERE score_id = $5
`

type UpdateScoreParams struct {
	TestID         pgtype.Int4 `json:"test_id"`
	ReadingScore   pgtype.Int4 `json:"reading_score"`
	ListeningScore pgtype.Int4 `json:"listening_score"`
	TotalScore     pgtype.Int4 `json:"total_score"`
	ScoreID        int64       `json:"score_id"`
}

func (q *Queries) UpdateScore(ctx context.Context, arg UpdateScoreParams) error {
	_, err := q.db.Exec(ctx, updateScore,
		arg.TestID,
		arg.ReadingScore,
		arg.ListeningScore,
		arg.TotalScore,
		arg.ScoreID,
	)
	return err
}
