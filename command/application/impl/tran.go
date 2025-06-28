package impl

import (
	"command-service/command/infra/sqlboiler/handler"
	"context"
	"database/sql"
	"log"

	"github.com/aarondl/sqlboiler/v4/boil"
)

type transaction struct {
}

func (t *transaction) begin(ctx context.Context) (*sql.Tx, error) {
	tran, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, handler.DBErrHandler(err)
	}
	return tran, nil
}

func (t *transaction) complete(tran *sql.Tx, err error) error {
	if err != nil {
		if e := tran.Rollback(); e != nil {
			return handler.DBErrHandler(err)
		} else {
			log.Println("Transaction rolled back")
		}
	} else {
		if e := tran.Commit(); e != nil {
			return handler.DBErrHandler(err)
		} else {
			log.Println("Transaction committed successfully")
		}
	}
	return nil
}
