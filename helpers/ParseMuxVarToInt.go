package helpers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseMuxVarToInt(r *http.Request, v string) int {
	value := mux.Vars(r)[v]

	value_int, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("strconvError:", err)
	}

	return value_int

}
