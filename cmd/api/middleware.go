package main

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

func (app *application) rateLimiter(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if enabled then use limiter
		if app.config.limiter.enabled {
			// initialize limiter
			limiter := rate.NewLimiter(rate.Limit(app.config.limiter.burst), app.config.limiter.burst)
			// check and limit requests
			if !limiter.Allow() {
				app.rateLimitExceededResponse(w, r)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// Why recover?
// its better to let client know that something wrong happened with 500 error rather than dropping connection in middle of call
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a deferred function (which will always be run in the event of a panic
		// as Go unwinds the stack).
		defer func() {
			if err := recover(); err != nil {
				// If there was a panic, set a "Connection: close" header on the
				// response. This acts as a trigger to make Go's HTTP server
				// automatically close the current connection after a response has been
				// sent.
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
