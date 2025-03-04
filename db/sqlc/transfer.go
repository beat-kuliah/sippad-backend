package db

import (
	"context"
)

type TransferTxResponse struct {
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	EntryIn     Entry    `json:"entry_in"`
	EntryOut    Entry    `json:"entry_out"`
	Transfer    Transfer `json:"transfer"`
}

func (s *Store) TransferTx(c context.Context, tr CreateTransferParams) (TransferTxResponse, error) {
	var tx TransferTxResponse
	var errT error

	err := s.execTx(c, func(q *Queries) error {
		// Transaction money
		tx.Transfer, errT = q.CreateTransfer(context.Background(), tr)
		if errT != nil {
			return errT
		}

		// Record entries
		inEarg := CreateEntryParams{
			AccountID: tr.ToAccountID,
			Amount:    tr.Amount,
		}
		tx.EntryIn, errT = q.CreateEntry(context.Background(), inEarg)
		if errT != nil {
			return errT
		}

		outEarg := CreateEntryParams{
			AccountID: tr.FromAccountID,
			Amount:    -1 * tr.Amount,
		}
		tx.EntryOut, errT = q.CreateEntry(context.Background(), outEarg)
		if errT != nil {
			return errT
		}

		// Update balance
		toArg := UpdateAccountBalanceNewParams{
			Amount: tr.Amount,
			ID:     int64(tr.ToAccountID),
		}
		tx.ToAccount, errT = q.UpdateAccountBalanceNew(context.Background(), toArg)
		if errT != nil {
			return errT
		}
		fromArg := UpdateAccountBalanceNewParams{
			Amount: -1 * tr.Amount,
			ID:     int64(tr.FromAccountID),
		}
		tx.FromAccount, errT = q.UpdateAccountBalanceNew(context.Background(), fromArg)
		if errT != nil {
			return errT
		}

		return nil
	})

	return tx, err
}
