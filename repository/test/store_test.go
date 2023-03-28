package repository

import (
	"context"
	"testing"

	"github.com/junanda/simple-bank/entity"
	. "github.com/junanda/simple-bank/repository"
	"github.com/stretchr/testify/require"
)

func TestTransferTX(t *testing.T) {
	store := NewStoreRepository(testDB, testTranfer, testEntry)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run e concurrent transfer transactiom
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan entity.TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), entity.TransferTx{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// Check result
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// Check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountId)
		require.Equal(t, account2.ID, transfer.ToAccountId)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = testTranfer.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// Check Entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountId)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = testEntry.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountId)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = testEntry.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// TODO: Checks account balance

	}
}
