package helpers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/subramanyam-searce/product-catalog-go/constants/responses"
)

var db *sql.DB

type JSONResponse struct {
	Message string `json:"message"`
}

func init() {
	connection_string := "user=service-pc-api host=localhost password=pcapi sslmode=disable dbname=product-catalog"
	var err error
	db, err = sql.Open("postgres", connection_string)
	if err != nil {
		fmt.Println("sqlOpenError:", err)
	}
}

func ConnectToDB() *sql.DB {
	return db
}

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

func HandleTestError(errorString string, err error, t *testing.T) {
	if err != nil {
		t.Log(errorString+":", err)
	}
}

func ParseMuxVarToInt(r *http.Request, v string) int {
	value := mux.Vars(r)[v]

	value_int, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("strconvError:", err)
	}

	return value_int
}

func RunQuery(query string, v ...any) (*sql.Rows, error) {
	db := ConnectToDB()
	var err error
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(v...)
	stmt.Close()

	return rows, err
}

func SendResponse(v any, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func SendJSONResponse(message string, w http.ResponseWriter) {
	SendResponse(JSONResponse{Message: message}, w)
}

func Paginate(page_no, total_items, items_per_page int) (int, int, error) {
	t := total_items
	n := items_per_page

	max_page_no := (t / n)
	if t%n != 0 {
		max_page_no += 1
	}

	if page_no <= max_page_no {
		start_index := ((page_no - 1) * n)
		end_index := start_index + items_per_page

		if total_items < start_index+items_per_page {
			end_index = total_items
		}

		return start_index, end_index, nil
	} else {
		return max_page_no, max_page_no, errors.New(responses.ProductsOutOfRange)
	}
}
