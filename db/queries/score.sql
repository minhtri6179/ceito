-- name: CreateScore :one
INSERT INTO score (
        reading_score,
        listening_score,
        total_score
    )
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetScore :one
SELECT *
FROM score
WHERE id = $1;
-- name: ListScores :many
SELECT *
FROM score
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: UpdateScore :exec
UPDATE score
SET reading_score = $1,
    listening_score = $2,
    total_score = $3
WHERE id = $4;
-- name: DeleteScore :exec
DELETE FROM score
WHERE id = $1;