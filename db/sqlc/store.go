package db

import (
	"context"
	"database/sql"
	"fmt"
)

// #begin Tx
// Transfer money
// Enter entry 1 in
// Enter entry 2 out
// Update balance
// #commit transaction

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTx(c context.Context, fq func(q *Queries) error) error {
	// initialize transactionst
	tx, err := s.db.BeginTx(c, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fq(q)

	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return fmt.Errorf("Encountered rollback error: %v", txErr)
		}
		return err
	}

	return tx.Commit()
}
