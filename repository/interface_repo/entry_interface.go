package interfaceRepo

import (
	"context"

	"github.com/junanda/simple-bank/entity"
)

type EntryRepository interface {
	CreateEntry(ctx context.Context, arg entity.EntryCreate) (entity.Entry, error)
	GetEntry(ctx context.Context, id int64) (entity.Entry, error)
	ListEntries(ctx context.Context, arg entity.ListEntriesParams) ([]entity.Entry, error)
}
