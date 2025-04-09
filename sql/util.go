package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var ErrTxNotFound = fmt.Errorf("transaction value not found")

type txSessionKeyType int

const txSessionKey = 0

// ContextWithTx returns a new context.Context with a *sql.Tx value.
func ContextWithTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txSessionKey, tx)
}

// TxFromContext extracts a *sql.Tx from a given context.Context
func TxFromContext(ctx context.Context) (*sql.Tx, error) {
	tx, ok := ctx.Value(txSessionKey).(*sql.Tx)
	if !ok {
		return nil, ErrTxNotFound
	}

	return tx, nil
}

// Executor returns a *sql.Tx or *sqlx.DB to execute a DB command based on the given context.
func Executor(ctx context.Context, db *sqlx.DB) sqlx.ExecerContext {
	if tx, ok := ctx.Value(txSessionKey).(*sql.Tx); ok {
		return tx
	}

	return db
}
