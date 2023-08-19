package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joshuabl97/urlShort/data"
	"github.com/joshuabl97/urlShort/handler"
	"github.com/rs/zerolog"
)

var portNum = flag.String("port_number", "8080", "The port number the server runs on")
var yamlPath = flag.String("yaml_filepath", "", "File used to prepopulate DB")
var timeZone = flag.String("timezone", "Etc/Greenwich", "Timezone formatted like the example provided")

func main() {
	// parse flags
	flag.Parse()

	// instantiate logger
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	// setting timezone
	loc, err := time.LoadLocation(*timeZone)
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
	db, _ = data.AddEndpoint(&l, db, "example1", "https://google.com")
	db, _ = data.AddEndpoint(&l, db, "example2", "https://example.com/")

	// prepopulate db
	if *yamlPath != "" {
		endpoints, err := data.ParseYaml(*yamlPath, db)
		if err != nil {
			l.Error().Err(err).Msg("Error parsing yaml")
		}
		db = data.AddMultipleEndpoints(endpoints, &l, db)
	}

	// test to see if endpoints were generated in db
	_, err = data.GetEndpoints(&l, db)
	if err != nil {
		l.Fatal().Err(err).Msg("Failed to initialize DB")
		os.Exit(1)
	}

	// helper handler contains *zerolog.Logger and *sql.DB
	hh := handler.NewHandlerHelper(&l, db)

	// registering the handlers on the serve mux (sm)
	sm := chi.NewRouter()
	sm.Get("/{endpoint}", handler.Logger(hh.Redirect, &l))
	sm.Get("/shortcuts", handler.Logger(hh.GetShortcuts, &l))
	sm.Post("/shortcut", handler.Logger(hh.CreateShortcut, &l))
	sm.Put("/{endpoint}", handler.Logger(hh.UpdateEndpoint, &l))
	sm.Delete("/{endpoint}", handler.Logger(hh.DeleteEndpoint, &l))

	// create a new server
	s := http.Server{
		Addr:         ":" + *portNum,           // configure the bind address
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
		l.Fatal().Err(err).Msg("Shutdown exceeded timeout")
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
