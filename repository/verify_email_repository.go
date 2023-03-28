package repository

import (
	"context"

	"github.com/junanda/simple-bank/db"
	"github.com/junanda/simple-bank/entity"
	interfaceRepo "github.com/junanda/simple-bank/repository/interface_repo"
)

func NewVerifyEmailRepository(db db.DBTX) interfaceRepo.VerifyEmailRepository {
	return &VerifyEmailRepositoryImpl{
		db: db,
	}
}

type VerifyEmailRepositoryImpl struct {
	db db.DBTX
}

func (v *VerifyEmailRepositoryImpl) CreateVerifyEmail(ctx context.Context, arg entity.CreateVerifyEmail) (entity.VerifyEmail, error) {
	var result entity.VerifyEmail

	query := `INSERT INTO verify_emails (username, email, secret_code) VALUES ($1, $2, $3) RETURNING id, username, email, secret_code, is_used, created_at, expired_at`
	row := v.db.QueryRowContext(ctx, query, arg.Username, arg.Email, arg.SecretCode)
	err := row.Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.SecretCode,
		&result.IsUsed,
		&result.CreatedAt,
		&result.ExpiredAt,
	)
	return result, err
}

func (v *VerifyEmailRepositoryImpl) UpdateVerifyEmail(ctx context.Context, arg entity.UpdateVerifyEmail) (entity.VerifyEmail, error) {
	var result entity.VerifyEmail
	queryUpdate := `UPDATE verify_emails SET is_used = TRUE WHERE id=$1 AND secret_code=$2 AND is_used=FALSE AND expired_at > now()
					RETURNING id, username, email, secret_code, is_used, created_at, expired_at`
	row := v.db.QueryRowContext(ctx, queryUpdate, arg.ID, arg.SecretCode)
	err := row.Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.SecretCode,
		&result.IsUsed,
		&result.CreatedAt,
		&result.ExpiredAt,
	)
	return result, err
}
