package helpers

import (
	"fmt"
)

func HandleError(errorString string, err error) {
	if err != nil {
		fmt.Println(errorString+":", err)
	}
}
