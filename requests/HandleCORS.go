package requests

import (
	"github.com/Infinitare/types-template/metrics"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
	"strings"
)

func HandleCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.String(), "BYsDXTr0gEBWK35kQf0i") {
			w.WriteHeader(http.StatusOK)
			return
		}

		traceID, spanID := uint64(0), uint64(0)
		span, ok := tracer.SpanFromContext(r.Context())
		if ok {
			traceID, spanID = span.Context().TraceID(), span.Context().SpanID()
		}

		metrics.Request(r)

		w.Header().Set("Access-Control-Allow-Origin", Domains.Website)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Datadog-Origin, X-Datadog-Parent-Id, X-Datadog-Sampling-Priority, X-Datadog-Trace-Id, X-Datadog-Span-Id, Traceparent, Process")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions && r.Header.Get("sec-Fetch-Mode") != "no-cors" {
			w.WriteHeader(http.StatusOK)
			return
		}

		defer HandlePanic(w, traceID, spanID)
		h.ServeHTTP(w, r)
	})
}
