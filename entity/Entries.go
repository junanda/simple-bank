package entity

import "time"

type Entry struct {
	ID        int64
	AccountId int64
	Amount    int64
	CreatedAt time.Time
}

type EntryCreate struct {
	AccountID int64
	Amount    int64
}

type ListEntriesParams struct {
	AccountID int64
	Limit     int32
	Offset    int32
}
