package handler

import (
	"command-service/internal/errs"
	"errors"
	"log"
	"net"

	"github.com/go-sql-driver/mysql"
)

func DBErrHandler(err error) error {

	var (
		opErr     *net.OpError
		driverErr *mysql.MySQLError
	)

	if errors.As(err, &opErr) {
		log.Println(err.Error())
		return errs.NewInternalError(opErr.Error())
	} else if errors.As(err, &driverErr) {
		log.Printf("Code: %d, Message: %s\n", driverErr.Number, driverErr.Message)
		if driverErr.Number == 1062 {
			return errs.NewCRUDError("duplicate entry error")
		} else {
			return errs.NewInternalError(driverErr.Message)
		}
	} else {
		log.Println(err.Error())
		return errs.NewInternalError(err.Error())
	}
}
