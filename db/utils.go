package db

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/leonardogazio/golang-account-crud/utils"
	"go.uber.org/zap"
)

// Persist ...
func Persist(db *sqlx.DB, sql string, args ...any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx := db.MustBegin().Tx
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, sql, args...); err != nil {
		utils.Logger.Error(strPersistFail, zap.Error(err))
		return err
	}

	if err := tx.Commit(); err != nil {
		utils.Logger.Error(strPersistFail, zap.Error(err))
		return err
	}

	return nil
}
