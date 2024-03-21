-- name: CreateUser :one
INSERT INTO users (username, email, full_name, password_hashed)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1;
-- name: UpdateUser :exec
UPDATE users
SET username = $1,
    email = $2,
    full_name = $3,
    password_hashed = $4
WHERE id = $4;
-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;