package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/watariRyo/tasktree/server/config"
)

type DBConnectionManager interface {
	Open(c *config.Config) (DBConnection, error)
}

type DBConnection interface {
	boil.ContextExecutor
}

type DBTransaction func(ctx context.Context, conn DBConnection, txFunc func(ctx context.Context, conn DBConnection) error, options *sql.TxOptions) error

func IsNoRecordError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
