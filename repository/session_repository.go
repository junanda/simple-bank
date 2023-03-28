package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/junanda/simple-bank/db"
	"github.com/junanda/simple-bank/entity"
	interfaceRepo "github.com/junanda/simple-bank/repository/interface_repo"
)

func NewSessionRepository(db db.DBTX) interfaceRepo.SessionRepository {
	return &SessionRepositoryImpl{
		db: db,
	}
}

type SessionRepositoryImpl struct {
	db db.DBTX
}

func (s *SessionRepositoryImpl) CreateSession(ctx context.Context, arg entity.SessionParams) (entity.Session, error) {
	var result entity.Session
	query := `INSERT INTO sessions (
		id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7
	) RETURNING id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at`

	row := s.db.QueryRowContext(ctx, query, arg.ID, arg.Username, arg.RefreshToken, arg.UserAgent, arg.ClientIp, arg.IsBlocked, arg.ExpiresAt)
	err := row.Scan(
		&result.ID,
		&result.Username,
		&result.RefreshToken,
		&result.UserAgent,
		&result.ClientIp,
		&result.IsBlocked,
		&result.ExpiresAt,
	)

	return result, err
}

func (s *SessionRepositoryImpl) GetSession(ctx context.Context, id uuid.UUID) (entity.Session, error) {
	var result entity.Session

	queryGet := `SELECT id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at FROM sessions WHERE id = $1 LIMIT 1`

	row := s.db.QueryRowContext(ctx, queryGet, id)
	err := row.Scan(
		&result.ID,
		&result.Username,
		&result.RefreshToken,
		&result.UserAgent,
		&result.ClientIp,
		&result.IsBlocked,
		&result.ExpiresAt,
	)

	return result, err
}
