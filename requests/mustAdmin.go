package requests

import (
	"context"
	"errors"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
	"os"
	"types-template/console"
	"types-template/helper"
)

func MustAdmin(handle func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trace, ctx := tracer.StartSpanFromContext(r.Context(), "admin", tracer.Measured(), tracer.ResourceName(helper.GetFunctionName(MustAdmin)))
		defer trace.Finish()

		admin := r.Header.Get("admin")
		if admin != os.Getenv("ADMIN_PASSWORD") {
			console.ErrorResponse(w, r, errors.New(http.StatusText(http.StatusUnauthorized)), http.StatusUnauthorized)
			return
		}

		*r = *r.WithContext(context.WithValue(ctx, "admin", true))
		handle(w, r)
	}
}
