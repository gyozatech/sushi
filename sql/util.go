package sql

import (
	"context"
	"database/sql"
	"fmt"
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

// DB interface contains the common methods between sql.DB and sql.Tx objects
type DB interface {
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

// Executor returns a DbOrTx (either *sql.Tx or *sql.DB) to execute a DB command based on the given context.
func Executor(ctx context.Context, db *sql.DB) DB {
	if tx, ok := ctx.Value(txSessionKey).(*sql.Tx); ok {
		return tx
	}
	return db
}
