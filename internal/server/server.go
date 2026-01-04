package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer(ctx context.Context, server *http.Server, shutdownTimeout time.Duration) error {
	serverErr := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
		close(serverErr)
	}()

	signalErr := make(chan os.Signal, 1)

	signal.Notify(signalErr, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-serverErr:
		return err

	case <-signalErr:
		log.Println("Signal Error Recived Closing Server")

	case <-ctx.Done():
		log.Println("Context closed")
	}

	shutdownContext, cancel := context.WithTimeout(
		context.Background(),
		shutdownTimeout,
	)

	defer cancel()

	if err := server.Shutdown(shutdownContext); err != nil {
		if closeErr := server.Close(); closeErr != nil {
			return errors.Join(err, closeErr)
		}
		return err
	}

	log.Println("Server Closed Gracefully")
	return nil

}
