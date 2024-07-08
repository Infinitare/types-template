package metrics

import (
	"github.com/Infinitare/mannaq-types/helper"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
)

func TraceHandler(handle func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trace, _ := tracer.StartSpanFromContext(r.Context(), "handler", tracer.Measured(), tracer.ResourceName(helper.GetFunctionName(handle)))
		handle(w, r)
		trace.Finish()
	}
}
