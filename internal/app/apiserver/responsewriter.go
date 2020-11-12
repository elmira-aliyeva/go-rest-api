package apiserver

import "net/http"

type responsewriter struct {
	http.ResponseWriter
	code int
}

func (w *responsewriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
