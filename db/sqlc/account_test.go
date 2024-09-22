package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "Vansh",
		Balance:  100,
		Currency: "INR",
	}
	acc, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)
}

func TestListAccounts(t *testing.T) {
	arg := ListAccountsParams{
		Limit: 10,
        Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(),arg)
    require.NoError(t, err)
    require.NotEmpty(t, accounts)
}
