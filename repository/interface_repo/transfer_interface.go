package interfaceRepo

import (
	"context"

	"github.com/junanda/simple-bank/entity"
)

type TransferRepository interface {
	CreateTransfer(ctx context.Context, arg entity.TransferTx) (entity.Transfer, error)
	GetTransfer(ctx context.Context, id int64) (entity.Transfer, error)
	ListTransfers(ctx context.Context, arg entity.ListTransferTx) ([]entity.Transfer, error)
}
