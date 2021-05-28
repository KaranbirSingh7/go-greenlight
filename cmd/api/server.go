package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// start a new goroutine in bg that will listen for SIG signals
	go func() {
		// create channel to hold any signals
		quit := make(chan os.Signal, 1)

		// listen for any signals from system
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// check for signal
		s := <-quit

		// log message about quitting
		app.logger.PrintInfo("caught signal", map[string]string{
			"signal": s.String(),
		})

		// quit server
		os.Exit(0)
	}()
	// Likewise log a "starting server" message.
	app.logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  app.config.env,
	})

	// Start the server as normal, returning any error.
	return srv.ListenAndServe()
}
