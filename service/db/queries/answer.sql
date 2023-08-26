-- name: CreateAnswer :one
INSERT INTO answer (question_id, answer_text, is_correct)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetAnswer :one
SELECT *
FROM answer
WHERE answer_id = $1;
-- name: ListAnswers :many
SELECT *
FROM answer
ORDER BY answer_id
LIMIT $1 OFFSET $2;
-- name: UpdateAnswer :exec
UPDATE answer
SET question_id = $1,
    answer_text = $2,
    is_correct = $3
WHERE answer_id = $4;
-- name: DeleteAnswer :exec
DELETE FROM answer
WHERE answer_id = $1;