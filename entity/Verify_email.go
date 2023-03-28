package entity

import "time"

type VerifyEmail struct {
	ID         int64
	Username   string
	Email      string
	SecretCode string
	IsUsed     bool
	CreatedAt  time.Time
	ExpiredAt  time.Time
}

type CreateVerifyEmail struct {
	Username   string
	Email      string
	SecretCode string
}

type UpdateVerifyEmail struct {
	ID         int64
	SecretCode string
}
