package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()

	for _, v := range Routes {
		router.HandleFunc(v.Path, v.Handler).Methods(v.Method)
	}

	fmt.Println("Server started on PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
