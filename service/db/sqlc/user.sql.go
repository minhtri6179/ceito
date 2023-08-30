// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, email, full_name, password_hashed)
VALUES ($1, $2, $3, $4)
RETURNING username, email, full_name, password_hashed, created_at, update_at
`

type CreateUserParams struct {
	Username       string      `json:"username"`
	Email          pgtype.Text `json:"email"`
	FullName       pgtype.Text `json:"full_name"`
	PasswordHashed pgtype.Text `json:"password_hashed"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.FullName,
		arg.PasswordHashed,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.FullName,
		&i.PasswordHashed,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	_, err := q.db.Exec(ctx, deleteUser, username)
	return err
}

const getUser = `-- name: GetUser :one
SELECT username, email, full_name, password_hashed, created_at, update_at
FROM users
WHERE username = $1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.FullName,
		&i.PasswordHashed,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET username = $1,
    email = $2,
    full_name = $3,
    password_hashed = $4
WHERE id = $4
`

type UpdateUserParams struct {
	Username       string      `json:"username"`
	Email          pgtype.Text `json:"email"`
	FullName       pgtype.Text `json:"full_name"`
	PasswordHashed pgtype.Text `json:"password_hashed"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.Username,
		arg.Email,
		arg.FullName,
		arg.PasswordHashed,
	)
	return err
}
