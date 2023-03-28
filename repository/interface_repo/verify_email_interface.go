package interfaceRepo

import (
	"context"

	"github.com/junanda/simple-bank/entity"
)

type VerifyEmailRepository interface {
	CreateVerifyEmail(ctx context.Context, arg entity.CreateVerifyEmail) (entity.VerifyEmail, error)
	UpdateVerifyEmail(ctx context.Context, arg entity.UpdateVerifyEmail) (entity.VerifyEmail, error)
}
