package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Username          string
	HashedPassword    string
	FullName          string
	Email             string
	PasswordChangedAt time.Time
	CreatedAt         time.Time
	IsEmailVerified   bool
}

type CreateUserParams struct {
	Username       string
	HashedPassword string
	FullName       string
	Email          string
}

type UpdateUser struct {
	HashedPassword    sql.NullString
	PasswordChangedAt sql.NullTime
	FullName          sql.NullString
	Email             sql.NullString
	IsEmailVerified   sql.NullBool
	Username          string
}
