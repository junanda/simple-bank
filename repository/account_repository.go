package repository

import (
	"context"
	"database/sql"

	"github.com/junanda/simple-bank/db"
	"github.com/junanda/simple-bank/entity"
	ifc "github.com/junanda/simple-bank/repository/interface_repo"
)

func NewAccountRepository(db db.DBTX) ifc.AccountRepository {
	return &AccountRepositoryImpl{
		db: db,
	}
}

type AccountRepositoryImpl struct {
	db db.DBTX
}

func (a *AccountRepositoryImpl) WithTx(tx *sql.Tx) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{
		db: tx,
	}
}

func (a *AccountRepositoryImpl) CreateAccount(ctx context.Context, arg entity.CreateAccountParams) (entity.Account, error) {
	var (
		i   entity.Account
		err error
	)

	query := `INSERT INTO accounts (
		owner,
		balance,
		currency
	) VALUES (
		$1, $2, $3
	) RETURNING id, owner, balance, currency, created_at`

	row := a.db.QueryRowContext(ctx, query, arg.Owner, arg.Balance, arg.Currency)

	err = row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}
func (a *AccountRepositoryImpl) GetAccount(ctx context.Context, id int64) (entity.Account, error) {
	var (
		data  entity.Account
		err   error
		query string
	)

	query = `SELECT id, owner, balance, currency, created_at FROM accounts WHERE id = $1 LIMIT 1`
	row := a.db.QueryRowContext(ctx, query, id)

	err = row.Scan(
		&data.ID,
		&data.Owner,
		&data.Balance,
		&data.Currency,
		&data.CreatedAt,
	)

	return data, err
}

func (a *AccountRepositoryImpl) ListAccounts(ctx context.Context, arg entity.ListAccountParams) ([]entity.Account, error) {
	var (
		items []entity.Account
		i     entity.Account
		err   error
		query string
	)

	query = `SELECT id, owner, balance, currency, created_at FROM accounts ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := a.db.QueryContext(ctx, query, arg.Limit, arg.Offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&i.ID, &i.Owner, &i.Balance, &i.Currency, &i.CreatedAt); err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, err
}

func (a *AccountRepositoryImpl) UpdateAccount(ctx context.Context, arg entity.UpdateAccountParams) (entity.Account, error) {
	var (
		data  entity.Account
		err   error
		query string
	)

	query = `UPDATE accounts SET balance = $2 WHERE id = $1
	RETURNING id, owner, balance, currency, created_at`

	row := a.db.QueryRowContext(ctx, query, arg.ID, arg.Balance)
	err = row.Scan(
		&data.ID,
		&data.Owner,
		&data.Balance,
		&data.Currency,
		&data.CreatedAt,
	)

	return data, err
}

func (a *AccountRepositoryImpl) DeleteAccount(ctx context.Context, id int64) error {
	queryDelete := `DELETE FROM accounts WHERE id = $1`
	_, err := a.db.ExecContext(ctx, queryDelete, id)
	return err
}
