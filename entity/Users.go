package entity

import "time"

type User struct {
	Username          string
	HashedPassword    string
	FullName          string
	Email             string
	PasswordChangedAt time.Time
	CreatedAt         time.Time
	IsEmailVerified   bool
}
