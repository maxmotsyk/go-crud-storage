package logging

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	// Якщо статус ще не встановлений, вважаємо його 200
	if lrw.statusCode == 0 {
		lrw.statusCode = http.StatusOK
	}
	return lrw.ResponseWriter.Write(b)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &loggingResponseWriter{ResponseWriter: w}

		start := time.Now()
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)

		log.WithFields(log.Fields{
			"method":   r.Method,
			"status":   lrw.statusCode,
			"duration": duration,
			"agent":    r.UserAgent(),
			"ip":       r.RemoteAddr,
			"path":     r.URL.Path,
		}).Info("Handled request")
	})
}
