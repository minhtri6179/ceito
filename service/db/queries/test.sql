-- name: CreateTest :one
INSERT INTO test (username)
VALUES ($1)
RETURNING *;
-- name: GetTest :one
SELECT *
FROM test
WHERE test_id = $1;
-- name: ListTests :many
SELECT *
FROM test
ORDER BY test_id
LIMIT $1 OFFSET $2;
-- name: UpdateTest :exec
UPDATE test
SET username = $1
WHERE test_id = $2;
-- name: DeleteTest :exec
DELETE FROM test
WHERE test_id = $1;