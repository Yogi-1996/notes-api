package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Yogi-1996/notes-backend/internal/config"
	"github.com/Yogi-1996/notes-backend/internal/database"
	"github.com/Yogi-1996/notes-backend/internal/handlers"
	"github.com/Yogi-1996/notes-backend/internal/middelware"
	"github.com/Yogi-1996/notes-backend/internal/repository"
	"github.com/Yogi-1996/notes-backend/internal/servers"
	"github.com/Yogi-1996/notes-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Load()
	db, err := database.NewPostgres(config)

	if err != nil {
		log.Fatal("error opening Database %w", err)
	}

	notesdb := repository.NewNoteRepositry(db)
	noteservice := services.NewNoteService(notesdb)
	notehandler := handlers.NewNoteHandler(noteservice)

	usersdb := repository.NewUserRepository(db)
	userservice := services.NewUserService(usersdb)
	userhandler := handlers.NewUserHandler(userservice)

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Note Server Running",
		})
	})

	api := router.Group("api")

	notes := api.Group("/", middelware.AunthMiddelware)
	{
		notes.POST("notes", notehandler.NoteAdd)
		notes.GET("notes", notehandler.GetNote)
		notes.GET("notes/:id", notehandler.GetNotesByID)
		notes.PUT("notes/:id", notehandler.ModNote)
		notes.DELETE("notes/:id", notehandler.DelNote)
	}

	auth := api.Group("/auth")
	{
		auth.POST("/register", userhandler.UserRegister)
		auth.POST("/login", userhandler.UserLogin)
	}

	ctx := context.Background()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := servers.RunServer(server, ctx, 10*time.Second); err != nil {
		log.Fatal(err)
	}
}
