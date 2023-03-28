package interfaceRepo

import (
	"context"

	"github.com/junanda/simple-bank/entity"
)

type Querier interface {
	CreateTransfer(ctx context.Context, arg entity.TransferTx) (entity.Transfer, error)
}

// var _ Querier = (*db.Queries)(nil)
