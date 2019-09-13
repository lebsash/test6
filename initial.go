package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"bitbucket.org/Sanny_Lebedev/test6/handlers"
	"bitbucket.org/Sanny_Lebedev/test6/logger"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/ssh/terminal"
)

func initial() *Config {
	var err error
	var file io.WriteCloser
	var log logger.Logger

	if terminal.IsTerminal(int(os.Stdout.Fd())) {

		log = logger.Logger{
			Logger: zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger(),
		}
	} else {
		deflogpath := "/var/log/"
		file, err = os.OpenFile(deflogpath+"test7.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Sprintf("can't write log file: %v", err))
		}

		log = logger.Logger{
			Logger: zerolog.New(file).With().Timestamp().Logger(),
		}
	}

	router := handlers.Router()

	log.Info().Str("Initial status", "OK").Time("Time", time.Now()).Msg("System init")

	return &Config{
		Log:       log,
		file:      file,
		router:    router,
		interrupt: make(chan os.Signal, 1),
		srv: &http.Server{
			Addr:    ":8000",
			Handler: router,
		},
	}
}
