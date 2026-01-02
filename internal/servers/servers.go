package servers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"
)

func RunServer(server *http.Server, ctx context.Context, timeout time.Duration) error {
	serverError := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverError <- err
		}
	}()

	signalError := make(chan os.Signal, 1)
	signal.Notify(signalError, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverError:
		return err

	case <-signalError:
		log.Println("shutdown signal recived")

	case <-ctx.Done():
		log.Println("Context Cancelled")
	}

	shutdownctx, cancel := context.WithTimeout(
		ctx,
		timeout,
	)

	defer cancel()

	if shutdownErr := server.Shutdown(shutdownctx); shutdownErr != nil {
		if closeErr := server.Close(); closeErr != nil {
			return errors.Join(shutdownErr, closeErr)
		}
		return shutdownErr

	}

	return nil

}
