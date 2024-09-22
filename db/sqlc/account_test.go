package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	
	acc, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// Create a new account
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	// Create a new account
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}
	// Update the account's balance
	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	require.Equal(t, account.CreatedAt, updatedAccount.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	// Create a new account
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
