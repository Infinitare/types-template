package requests

import (
	"github.com/pkg/errors"
	"net/http"
	"types-template/console"
)

func HandlePanic(w http.ResponseWriter, traceID, spanID uint64) {

	rec := recover()
	if rec != nil {
		var err error
		switch t := rec.(type) {
		case string:
			err = errors.New(t)
		case error:
			err = errors.Wrap(t, "panic")
		default:
			err = errors.New("unknown error")
		}

		console.RequestFatal(traceID, spanID, err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal Server Error"))
	}

}
