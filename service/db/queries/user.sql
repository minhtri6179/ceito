-- name: CreateUser :one
INSERT INTO users (username, email, full_name, password_hashed)
VALUES ($1, $2, $3, $4)
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
    full_name = $3,
    password_hashed = $4
WHERE id = $4;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;