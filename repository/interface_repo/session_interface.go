package interfaceRepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/junanda/simple-bank/entity"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, arg entity.SessionParams) (entity.Session, error)
	GetSession(ctx context.Context, id uuid.UUID) (entity.Session, error)
}
