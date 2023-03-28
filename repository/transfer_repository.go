package repository

import (
	"context"

	"github.com/junanda/simple-bank/db"
	"github.com/junanda/simple-bank/entity"
	ifc "github.com/junanda/simple-bank/repository/interface_repo"
)

func NewTransferImpl(db db.DBTX) ifc.TransferRepository {
	return &TransferRepositoryImpl{
		db: db,
	}
}

type TransferRepositoryImpl struct {
	db db.DBTX
}

func (q *TransferRepositoryImpl) CreateTransfer(ctx context.Context, arg entity.TransferTx) (entity.Transfer, error) {
	var (
		result entity.Transfer
	)

	query := `INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES ($1, $2, $3) RETURNING id, from_account_id, to_account_id, amount, created_at`
	row := q.db.QueryRowContext(ctx, query, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	err := row.Scan(
		&result.ID,
		&result.FromAccountId,
		&result.ToAccountId,
		&result.Amount,
		&result.CreatedAt,
	)

	return result, err
}

func (t *TransferRepositoryImpl) GetTransfer(ctx context.Context, id int64) (entity.Transfer, error) {
	var result entity.Transfer
	query := `SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers WHERE id = $1 LIMIT 1`

	row := t.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&result.ID,
		&result.FromAccountId,
		&result.ToAccountId,
		&result.Amount,
		&result.CreatedAt,
	)
	return result, err
}

func (t *TransferRepositoryImpl) ListTransfers(ctx context.Context, arg entity.ListTransferTx) ([]entity.Transfer, error) {
	var (
		result      []entity.Transfer
		tmpTransfer entity.Transfer
	)

	query := `SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers WHERE from_account_id = $1 OR to_account_id = $2 ORDER BY id LIMIT $3 OFFSET $4`

	row, err := t.db.QueryContext(ctx, query, arg.FromAccountID, arg.ToAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		if err = row.Scan(
			&tmpTransfer.ID,
			&tmpTransfer.FromAccountId,
			&tmpTransfer.ToAccountId,
			&tmpTransfer.Amount,
			&tmpTransfer.CreatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, tmpTransfer)
	}

	if err := row.Close(); err != nil {
		return nil, err
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
