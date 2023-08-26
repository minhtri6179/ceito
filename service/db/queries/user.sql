-- name: CreateUser :one
INSERT INTO users (username, email, password_hashed)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: UpdateUser :exec
UPDATE users
SET username = $1,
    email = $2,
    password_hashed = $3
WHERE id = $4;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;