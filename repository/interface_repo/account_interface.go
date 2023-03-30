package interfaceRepo

import (
	"context"

	"github.com/junanda/simple-bank/entity"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, arg entity.CreateAccountParams) (entity.Account, error)
	GetAccount(ctx context.Context, id int64) (entity.Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (entity.Account, error)
	ListAccounts(ctx context.Context, arg entity.ListAccountParams) ([]entity.Account, error)
	UpdateAccount(ctx context.Context, arg entity.UpdateAccountParams) (entity.Account, error)
	DeleteAccount(ctx context.Context, id int64) error
}
