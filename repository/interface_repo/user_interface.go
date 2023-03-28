package interfaceRepo

import (
	"context"

	"github.com/junanda/simple-bank/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg entity.CreateUserParams) (entity.User, error)
	GetUser(ctx context.Context, username string) (entity.User, error)
	UpdateUser(ctx context.Context, arg entity.UpdateUser) (entity.User, error)
}
