package main

import (
	"fmt"
	"os"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func main() {
	db := helpers.ConnectToDB()

	sqlString, err := os.ReadFile("sql_commands/testing_state.sql")
	helpers.HandleError("readFileError", err)
	if err != nil {
		fmt.Println("Failed to restore DB")
		return
	}

	_, err = db.Exec(string(sqlString))
	helpers.HandleError("dbExecError", err)
	if err != nil {
		fmt.Println("Failed to restore DB")
		return
	}

	fmt.Println("Successfully restored DB to testing state")
}
