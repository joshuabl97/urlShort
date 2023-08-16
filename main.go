package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joshuabl97/urlShort/data"
	"github.com/rs/zerolog"
)

func main() {
	fmt.Println("test")

	// instantiate logger
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	// setting timezone
	loc, err := time.LoadLocation("Etc/Greenwich")
	if err != nil {
		l.Error().Msg("Couldn't determine timezone, using local machine time")
	} else if err == nil {
		time.Local = loc
	}

	// make the logs look pretty
	l = l.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	// create a custom logger that wraps the zerolog.Logger we instantiated/customized above
	errorLog := &zerologLogger{l}

	// instantiate sqlite db
	db, err := data.MakeDb(&l)
	if err != nil {
		l.Fatal().Err(err).Msg("Error Creating DB... ")
		os.Exit(1)
	}
	defer db.Close()

	// add endpoints to db
	db, _ = data.AddToDB(&l, db, "example1", "https://wwww.google.com")

	// test to see if endpoints were generated in db
	_ = data.GetEndpoints(&l, db)

	// registering the handlers on the serve mux (sm)
	sm := chi.NewRouter()

	// create a new server
	s := http.Server{
		Addr:         ":8080",                  // configure the bind address
		Handler:      sm,                       // set the default handler
		IdleTimeout:  120 * time.Second,        // max duration to wait for the next request when keep-alives are enabled
		ReadTimeout:  5 * time.Second,          // max duration for reading the request
		WriteTimeout: 10 * time.Second,         // max duration before returning the request
		ErrorLog:     log.New(errorLog, "", 0), // set the logger for the server
	}

	// this go function starts the server
	// when the function is done running, that means we need to shutdown the server
	// we can do this by killing the program, but if there are requests being processed
	// we want to give them time to complete
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal().Err(err)
		}
	}()

	// sending kill and interrupt signals to os.Signal channel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// does not invoke 'graceful shutdown' unless the signalChannel is closed
	<-sigChan

	l.Info().Msg("Received terminate, graceful shutdown")

	// this timeoutContext allows the server 30 seconds to complete all requests (if any) before shutting down
	timeoutCtx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = s.Shutdown(timeoutCtx)
	if err != nil {
		l.Fatal().Msg("Shutdown exceeded timeout")
	}
}

// custom logger type that wraps zerolog.Logger
type zerologLogger struct {
	logger zerolog.Logger
}

// implement the io.Writer interface for our custom logger.
func (l *zerologLogger) Write(p []byte) (n int, err error) {
	l.logger.Error().Msg(string(p))
	return len(p), nil
}
