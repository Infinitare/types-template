package responses

import (
	"encoding/json"
	"github.com/Infinitare/types-template/console"
	"github.com/Infinitare/types-template/metrics"
	"net/http"
)

func SendJson(data interface{}, statusCode int, w http.ResponseWriter, r *http.Request) {

	if statusCode == http.StatusNoContent || statusCode == http.StatusNotModified {
		w.WriteHeader(statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		console.ErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}

	metrics.Response(r, data)

}
