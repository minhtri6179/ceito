-- name: CreateTest :one
INSERT INTO test_result (username)
VALUES ($1)
RETURNING *;
-- name: GetTest :one
SELECT *
FROM test_result
WHERE test_id = $1;
-- name: ListTests :many
SELECT *
FROM test_result
ORDER BY test_id
LIMIT $1 OFFSET $2;
-- name: UpdateTest :exec
UPDATE test_result
SET username = $1
WHERE test_id = $2;
-- name: DeleteTest :exec
DELETE FROM test_result
WHERE test_id = $1;