package helpers

import (
	"fmt"
	"log"
	"os"
)

func HandleError(errorString string, err error) {
	if err != nil {
		output := fmt.Sprint(errorString+":", err)
		fmt.Println(output)

		file, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		log.SetOutput(file)
		log.Println(output)
	}
}
