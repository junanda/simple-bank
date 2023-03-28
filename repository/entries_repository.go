package repository

import (
	"context"

	"github.com/junanda/simple-bank/db"
	"github.com/junanda/simple-bank/entity"
	interfaceRepo "github.com/junanda/simple-bank/repository/interface_repo"
)

type EntriesRepositoryImpl struct {
	db db.DBTX
}

func NewEntriesRepository(db db.DBTX) interfaceRepo.EntryRepository {
	return &EntriesRepositoryImpl{
		db: db,
	}
}

func (e *EntriesRepositoryImpl) CreateEntry(ctx context.Context, arg entity.EntryCreate) (entity.Entry, error) {
	var result entity.Entry

	query := `INSERT INTO entries (account_id, amount) VALUES ($1, $2) RETURNING id, account_id, amount, created_at`

	row := e.db.QueryRowContext(ctx, query, arg.AccountID, arg.Amount)
	err := row.Scan(
		&result.ID,
		&result.AccountId,
		&result.Amount,
		&result.CreatedAt,
	)

	return result, err
}

func (e *EntriesRepositoryImpl) GetEntry(ctx context.Context, id int64) (entity.Entry, error) {
	var result entity.Entry

	query := `SELECT id, account_id, amount, created_at FROM entries WHERE id = $1 LIMIT 1`

	row := e.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&result.ID,
		&result.AccountId,
		&result.Amount,
		&result.CreatedAt,
	)
	return result, err
}

func (e *EntriesRepositoryImpl) ListEntries(ctx context.Context, arg entity.ListEntriesParams) ([]entity.Entry, error) {
	var (
		results  []entity.Entry
		tmpEntry entity.Entry
	)

	queryList := `SELECT id, account_id, amount, crated_at FROM entries WHERE account_id = $1 LIMIT $2 OFFSET $3`
	rows, err := e.db.QueryContext(ctx, queryList, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&tmpEntry.ID,
			&tmpEntry.AccountId,
			&tmpEntry.Amount,
			&tmpEntry.CreatedAt,
		); err != nil {
			return nil, err
		}
		results = append(results, tmpEntry)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
