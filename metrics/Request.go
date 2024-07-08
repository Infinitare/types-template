package metrics

import (
	"bytes"
	"errors"
	"github.com/Infinitare/types-template/console"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"io"
	"net/http"
)

func Request(r *http.Request) {

	console.RequestInfo(r, r.Method, r.URL.Path)
	if r.Body == nil {
		return
	}

	trace, ok := tracer.SpanFromContext(r.Context())
	if !ok {
		console.RequestError(r, errors.New("no span found in context"))
		return
	}

	var body []byte
	if r.Body != nil {
		body, _ = io.ReadAll(r.Body)
		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	if len(body) == 0 {
		return
	}

	trace.SetTag("request", string(body))

}
