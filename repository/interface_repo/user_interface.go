package interfaceRepo

import (
	"context"

	"github.com/junanda/simple-bank/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg entity.CreateUserParams) (entity.User, error)
}
