package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}
// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountId int64   `json:"from_account_id"`
	ToAccountId   int64   `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}
// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}
// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
// execTx executes a function within a database transaction
func (store *Store) execTransaction(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	queries := New(tx)

	// generic function executed
	err = fn(queries)
	if err != nil {
		// in case rollback fails
		if rollBckError := tx.Rollback(); rollBckError != nil {
			return fmt.Errorf("transaction err: %v, rollback err %v", err, rollBckError)
		}
		return err
	}

	return tx.Commit()
}
// TransferTx performs a money transfer from one account to the other.
// It creates a transfer record, add account entries, and update accounts' balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTransaction(ctx, func(q *Queries) error {
		var err error

		// transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        strconv.FormatFloat(arg.Amount, 'f', 2, 64),
		})
		if err != nil {
			return err
		}

		// adding FROM account entries
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    strconv.FormatFloat(-arg.Amount, 'f', 2, 64),
		})
		if err != nil {
			return err
		}

		// adding TO account entries
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    strconv.FormatFloat(arg.Amount, 'f', 2, 64),
		})
		if err != nil {
			return err
		}

		//! TODO: update accounts' balance

		return nil
	})

	if err != nil {
		return result, err
	}

	return result, nil
}
