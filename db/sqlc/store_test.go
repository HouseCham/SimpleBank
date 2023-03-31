package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run n concurrent transfer transactions
	n := 10
	amount := int32(10)

	// we use chan keyword to obtain or recolect data from go routines
	errs := make(chan error)
	//results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: fromAccountID,
				ToAccountId:   toAccountID,
				Amount:        amount,
			})
			// we add erros and results from each go routine
			errs <- err
		}()
	}

	// check results
	//existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		
	}

	// check the final updated balances
	updatedAccoun1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccoun2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	//fmt.Println(">> after: ", account1.Balance, account2.Balance)
	require.Equal(t, account1.Balance, updatedAccoun1.Balance)
	require.Equal(t, account2.Balance, updatedAccoun2.Balance)
}
