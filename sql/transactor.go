package sql

import (
	"context"

	"github.com/jmoiron/sqlx"
)

/*
	Transactor interface allows you to bring transaction at Service layer and transact multiple different DAO calls committing them only if all
	of them are successful or else rollbacking them, even if these operations are spanned among multiple DAOs or even multiple services.
	In fact, the transaction is saved in the ctx and fetched and used when present.
*/

// Transactor is the interface to implement SQL transactions at the service layer level
type Transactor interface {
	WithinTransaction(ctx context.Context, txFunc func(txCtx context.Context) error) error
}

// TransactorImpl is the implementation of the Transactor interface
type TransactorImpl struct {
	// // Convert it to a *sqlx.DB = sqlx.NewDb(standardDB, "postgres")
	// vice versa: standardDB := xdb.DB
	db *sqlx.DB
}

// NewTransactor creates a new SQL transactor
func NewTransactor(db *sqlx.DB) Transactor {
	return &TransactorImpl{
		db: db,
	}
}

// WithinTransaction fetches the current transaction or creates a new one in which to perform all the calls
func (t *TransactorImpl) WithinTransaction(ctx context.Context, updateFn func(ctx context.Context) error) error {

	// let's initially assume the transaction comes from another caller, so this is a continuation
	scopedTransaction := false

	// if the context contains already a transaction, we use it as a part of a bigger transaction
	tx, err := TxFromContext(ctx)
	if err != nil {
		// no higher-level transaction found in the context: beginning a new transaction
		tx, err = t.db.Begin()
		if err != nil {
			return err
		}
		// the transaction should start and end in this function
		scopedTransaction = true
	}

	txContext := ContextWithTx(ctx, tx)

	if err := updateFn(txContext); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	// let's commit only if the transaction is scoped to this function, so if it shouldn't continue elsewhere
	if scopedTransaction {
		tx.Commit()
	}

	return nil
}
