-- name: CreateAccount :one
INSERT INTO accounts (owner, test_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1
LIMIT 1;
-- name: GetAccountForUpdate :one
SELECT *
FROM accounts
WHERE id = $1
LIMIT 1 FOR NO KEY
UPDATE;
-- name: ListAccounts :many
SELECT *
FROM accounts
WHERE owner = $1
ORDER BY id
LIMIT $2 OFFSET $3;
-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;