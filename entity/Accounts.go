package entity

import "time"

type Account struct {
	ID        int64
	Owner     string
	Balance   int64
	Currency  string
	CreatedAt time.Time
}

type CreateAccountParams struct {
	Owner    string
	Balance  int64
	Currency string
}

type ListAccountParams struct {
	Limit  int32
	Offset int32
}

type UpdateAccountParams struct {
	ID      int64
	Balance int64
}
