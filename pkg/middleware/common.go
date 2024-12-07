package middleware

import "net/http"

// WrapperWriter структура для получения статус кода
type WrapperWriter struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader функция получения статус кода
func (w *WrapperWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}
