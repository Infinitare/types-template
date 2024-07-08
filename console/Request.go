package console

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
	"reflect"
)

type StackTracer interface {
	StackTrace() errors.StackTrace
}

type errorField struct {
	Kind    string `json:"kind"`
	Stack   string `json:"stack"`
	Message string `json:"message"`
}

func newErrorField(err error) errorField {

	sErr, ok := err.(StackTracer)
	var st errors.StackTrace
	if ok {
		st = sErr.StackTrace()
	}

	stack := fmt.Sprintf("%+v", st)
	return errorField{
		Kind:    reflect.TypeOf(err).String(),
		Stack:   stack,
		Message: err.Error(),
	}

}

func ddFields(traceID, spanID uint64) logrus.Fields {
	return logrus.Fields{
		"trace_id": traceID,
		"span_id":  spanID,
	}
}

func requestLog(r *http.Request, err ...error) *logrus.Entry {
	span, ok := tracer.SpanFromContext(r.Context())
	if !ok {
		return logrus.WithFields(logrus.Fields{})
	}

	if len(err) > 0 {
		return errorLog(span.Context().TraceID(), span.Context().SpanID(), err[0])
	}

	return logger.WithFields(logrus.Fields{
		"dd": ddFields(span.Context().TraceID(), span.Context().SpanID()),
	})
}

func errorLog(traceID, spanID uint64, err error) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"error": newErrorField(err),
		"dd":    ddFields(spanID, traceID),
	})
}

func RequestInfo(r *http.Request, text ...any) {
	requestLog(r).Info(log(text...))
}

func RequestWarn(r *http.Request, text ...any) {
	requestLog(r).Warn(log(text...))
}

func RequestError(r *http.Request, err error) {
	requestLog(r, err).Error(log(err.Error()))
}

func RequestFatal(traceID, spanID uint64, err error) {
	errorLog(traceID, spanID, err).Error(log(err.Error()))
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, err error, code int) {
	if err == nil {
		err = errors.New("")
	}

	RequestError(r, err)
	http.Error(w, http.StatusText(code), code)
}
