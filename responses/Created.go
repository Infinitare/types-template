package responses

import (
	"net/http"
)

func SendCreated(w http.ResponseWriter, r *http.Request) {

	SendJson(Response{
		Status: http.StatusText(http.StatusCreated),
	}, http.StatusCreated, w, r)

}
