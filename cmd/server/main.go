package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Yogi-1996/notes-backend/internal/handlers"
	"github.com/Yogi-1996/notes-backend/internal/middelware"
	"github.com/Yogi-1996/notes-backend/internal/repository"
	"github.com/Yogi-1996/notes-backend/internal/servers"
	"github.com/Yogi-1996/notes-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {

	notesdb := repository.NewNoteRepositry()
	noteservice := services.NewNoteService(notesdb)
	notehandler := handlers.NewNoteHandler(noteservice)

	usersdb := repository.NewUserRepository()
	userservice := services.NewUserService(usersdb)
	userhandler := handlers.NewUserHandler(userservice)

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Note Server Running",
		})
	})

	api := router.Group("api")

	notes := api.Group("/notes", middelware.AunthMiddelware)
	{
		notes.POST("/", notehandler.NoteAdd)
		notes.GET("/", notehandler.GetNote)
		notes.GET("/:id", notehandler.GetNotesByID)
		notes.PUT("/:id", notehandler.ModNote)
		notes.DELETE("/:id", notehandler.DelNote)
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
