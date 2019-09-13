package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"syscall"

	"bitbucket.org/Sanny_Lebedev/test6/logger"
	"github.com/gorilla/mux"
)

// Config - configuration
type Config struct {
	Log       logger.Logger
	file      io.WriteCloser
	router    *mux.Router
	interrupt chan os.Signal
	srv       *http.Server
}

func main() {

	c := initial()
	defer c.file.Close()

	go func() {
		c.Log.Warn().Err(c.srv.ListenAndServe()).Msg("Test6")
	}()
	c.Log.Print("The service is ready to listen and serve.")

	c.Log.Warn().Err(http.ListenAndServe(c.srv.Addr, c.router)).Msg("Test6")

	killSignal := <-c.interrupt
	switch killSignal {
	case os.Interrupt:
		c.Log.Warn().Msg("Got SIGINT...")
	case syscall.SIGTERM:
		c.Log.Warn().Msg("Got SIGTERM...")
	}

	c.Log.Warn().Msg("The service is shutting down...")
	c.srv.Shutdown(context.Background())
	c.Log.Warn().Msg("TDone")

}
