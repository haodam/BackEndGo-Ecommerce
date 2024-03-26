// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user.sql

package repository

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
    ) RETURNING user_id, username, full_name, email, phone, avatar, country, city, street_address, zip, status, hashed_password, password_changed_at, created_at, updated_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.Avatar,
		&i.Country,
		&i.City,
		&i.StreetAddress,
		&i.Zip,
		&i.Status,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT user_id, username, full_name, email, phone, avatar, country, city, street_address, zip, status, hashed_password, password_changed_at, created_at, updated_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.Avatar,
		&i.Country,
		&i.City,
		&i.StreetAddress,
		&i.Zip,
		&i.Status,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
