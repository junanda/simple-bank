package entity

import "time"

type Entry struct {
	ID        int64
	AccountId int64
	Amount    int64
	CreatedAt time.Time
}
