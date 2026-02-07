package servers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Yogi-1996/notes-backend/internal/config"
	"github.com/Yogi-1996/notes-backend/internal/database"
	"github.com/Yogi-1996/notes-backend/internal/handlers"
	"github.com/Yogi-1996/notes-backend/internal/middelware"
	"github.com/Yogi-1996/notes-backend/internal/repository"
	"github.com/Yogi-1996/notes-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func RunApp(cfg *config.Config) error {

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	db, err := database.NewPostgres(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	
	noteRepo := repository.NewNoteRepositry(db)
	noteService := services.NewNoteService(noteRepo)
	noteHandler := handlers.NewNoteHandler(noteService)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	
	router := setupRouter(noteHandler, userHandler)

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	return runServer(ctx, server, 10*time.Second)
}

func setupRouter(noteHandler *handlers.NoteHandler, userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Note Server Running",
		})
	})

	api := router.Group("/api")

	notes := api.Group("/", middelware.AunthMiddelware)
	{
		notes.POST("notes", noteHandler.NoteAdd)
		notes.GET("notes", noteHandler.GetNote)
		notes.GET("notes/:id", noteHandler.GetNotesByID)
		notes.PUT("notes/:id", noteHandler.ModNote)
		notes.DELETE("notes/:id", noteHandler.DelNote)
	}

	auth := api.Group("/auth")
	{
		auth.POST("/register", userHandler.UserRegister)
		auth.POST("/login", userHandler.UserLogin)
	}

	return router
}

func runServer(ctx context.Context, server *http.Server, timeout time.Duration) error {
	serverErr := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		return err

	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		if closeErr := server.Close(); closeErr != nil {
			return errors.Join(err, closeErr)
		}
		return err
	}

	return nil
}
