-- name: CreateQuestion :one
INSERT INTO question (question_text, answer_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetQuestion :one
SELECT *
FROM question
WHERE question_id = $1;
-- name: ListQuestions :many
SELECT *
FROM question
ORDER BY question_id
LIMIT $1 OFFSET $2;
-- name: UpdateQuestion :exec
UPDATE question
SET question_text = $1,
    answer_id = $2
WHERE question_id = $3;
-- name: DeleteQuestion :exec
DELETE FROM question
WHERE question_id = $1;