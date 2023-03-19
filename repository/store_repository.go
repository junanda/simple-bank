package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/junanda/simple-bank/db"
)

func NewStoreRepository(dbS *sql.DB) *StoreRepositoryImpl {
	return &StoreRepositoryImpl{
		dbase:   dbS,
		Queries: db.New(dbS),
	}
}

type StoreRepositoryImpl struct {
	dbase *sql.DB
	*db.Queries
}

func (store *StoreRepositoryImpl) execTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := store.dbase.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
