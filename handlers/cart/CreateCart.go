package handlers_cart

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	ref := uuid.New()

	_, err := helpers.RunQuery("INSERT INTO cart_reference VALUES($1, $2);", ref, time.Now())
	helpers.HandleError("runQueryError:", err)

	if err != nil {
		helpers.SendResponse(map[string]string{"message": err.Error()}, w)
	} else {
		helpers.SendResponse(map[string]uuid.UUID{"ref": ref}, w)
	}

}
