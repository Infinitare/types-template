package responses

import (
	"net/http"
)

func SendOk(w http.ResponseWriter, r *http.Request) {

	SendJson(Response{
		Status: http.StatusText(http.StatusOK),
	}, http.StatusOK, w, r)

}
