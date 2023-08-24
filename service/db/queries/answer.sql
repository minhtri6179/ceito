-- name: createAnswer :one
INSERT INTO answer (question_id, answer_text, is_correct)
VALUES ($1, $2, $3)
RETURNING *;