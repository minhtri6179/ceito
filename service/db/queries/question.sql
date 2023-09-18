-- name: CreateQuestion :one
INSERT INTO question (
        question_text,
        answer_id,
        test_name,
        created_at
    )
VALUES ($1, $2, $3, $4)
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
    answer_id = $2,
    test_name = $3
WHERE question_id = $4;
-- name: DeleteQuestion :exec
DELETE FROM question
WHERE question_id = $1;