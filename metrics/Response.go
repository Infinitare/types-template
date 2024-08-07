package metrics

import (
	"encoding/json"
	"errors"
	"github.com/Infinitare/types-template/console"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
)

func Response(r *http.Request, data interface{}) {

	trace, ok := tracer.SpanFromContext(r.Context())
	if !ok {
		console.RequestError(r, errors.New("no span found in context"))
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		console.RequestError(r, err)
		return
	}

	trace.SetTag("response", string(jsonData))

}
