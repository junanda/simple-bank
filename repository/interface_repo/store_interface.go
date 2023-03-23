package interfaceRepo

import (
	"context"

	"github.com/junanda/simple-bank/entity"
)

type StoreRepository interface {
	TransferTx(ctx context.Context, arg entity.TransferTx) (entity.TransferTxResult, error)
}
