package entity

import "time"

type Transfer struct {
	ID            int64
	FromAccountId int64
	ToAccountId   int64
	Amount        int64
	CreatedAt     time.Time
}
