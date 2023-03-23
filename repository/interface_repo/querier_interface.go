package interfaceRepo

import "context"

type Querier interface {
	CreateTransfer(ctx context.Context)
}
