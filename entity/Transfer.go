package entity

import "time"

type Transfer struct {
	ID            int64
	FromAccountId int64
	ToAccountId   int64
	Amount        int64
	CreatedAt     time.Time
}

type TransferTx struct {
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
}

type TransferTxResult struct {
	Transfer    Transfer
	FromAccount Account
	ToAccount   Account
	FromEntry   Entry
	ToEntry     Entry
}

type ListTransferTx struct {
	FromAccountID int64
	ToAccountID   int64
	Limit         int32
	Offset        int32
}
