package repository

import (
	"context"
	"testing"

	"github.com/junanda/simple-bank/entity"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := entity.CreateAccountParams{
		Owner:    "junanda",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testAccounts.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
