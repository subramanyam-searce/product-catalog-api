package tests

import (
	"errors"
	"os"

	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func RestoreDBToTestingState() error {
	db := helpers.ConnectToDB()

	sqlString, err := os.ReadFile("../db/sql_commands/testing_state.sql")
	helpers.HandleError("readFileError", err)
	if err != nil {
		return errors.New(responses.FailedToRestoreDB)
	}

	_, err = db.Exec(string(sqlString))
	helpers.HandleError("dbExecError", err)
	if err != nil {
		return errors.New(responses.FailedToRestoreDB)
	}

	return nil
}
