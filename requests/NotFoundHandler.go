package requests

import (
	"github.com/Infinitare/types-template/console"
	"github.com/pkg/errors"
	"net/http"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	console.ErrorResponse(w, r, errors.New(http.StatusText(http.StatusNotFound)), http.StatusNotFound)
}

var NotFoundHandler = http.HandlerFunc(notFound)
