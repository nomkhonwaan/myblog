package log

import (
	"net/http"
	"time"
)

// Timer is a compatible interface for retrieving current system date-time
type Timer interface {
	// Return current system date-time
	Now() time.Time
}

// Outputer is a compatible interface for logging with format
type Outputer interface {
	// Log with format to the output
	Printf(format string, args ...interface{})
}

// LoggingInterceptor is an HTTP logger which logs in similar format as NGINX
type LoggingInterceptor struct {
	Timer
	Outputer
}

func (interceptor LoggingInterceptor) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := interceptor.Timer.Now()
		customResponseWriter := &CustomResponseWriter{ResponseWriter: w}

		h.ServeHTTP(customResponseWriter, r)

		interceptor.Outputer.Printf("\"%s %s\" %d %q %q \"%v\"",
			r.Method,
			r.RequestURI,
			customResponseWriter.statusCode,
			r.UserAgent(),
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

// CustomResponseWriter embeds http.ResponseWriter interface for overriding method WriteHeader
// which will save the response status code for logging purpose
type CustomResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// NewLoggingInterceptor creates new logger for logging HTTP request
func NewLoggingInterceptor(timer Timer, outputer Outputer) LoggingInterceptor {
	return LoggingInterceptor{Timer: timer, Outputer: outputer}
}
