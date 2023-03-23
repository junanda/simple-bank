package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/junanda/simple-bank/db"
	"github.com/junanda/simple-bank/entity"
	ifc "github.com/junanda/simple-bank/repository/interface_repo"
)

func NewStoreRepository(dbS *sql.DB) ifc.StoreRepository {
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

// Transfer transactiom perform a money transfer from one account to the other
// It creates a transfer  record, and account entries, and update accounts balance within  a single database transaction
func (store *StoreRepositoryImpl) TransferTx(ctx context.Context, arg entity.TransferTx) (entity.TransferTxResult, error) {
	var result entity.TransferTxResult

	err := store.execTx(ctx, func(q *db.Queries) error {
		var err error
		result, err = q.CreateTransfer(ctx, CreateTransferparams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CrateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CrateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		// TODO: update accounts balance

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
