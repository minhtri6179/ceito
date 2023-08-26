-- name: CreateScore :one
INSERT INTO score (
        test_id,
        reading_score,
        listening_score,
        total_score
    )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetScore :one
SELECT *
FROM score
WHERE score_id = $1;
-- name: ListScores :many
SELECT *
FROM score
ORDER BY score_id
LIMIT $1 OFFSET $2;
-- name: UpdateScore :exec
UPDATE score
SET test_id = $1,
    reading_score = $2,
    listening_score = $3,
    total_score = $4
WHERE score_id = $5;
-- name: DeleteScore :exec
DELETE FROM score
WHERE score_id = $1;