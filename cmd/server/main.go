package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Yogi-1996/notes-backend/internal/handlers"
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

	ginHandler := gin.Default()

	ginHandler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"messsage": "Note Server Running",
		})
	})

	ginHandler.POST("/notes", notehandler.NoteAdd)
	ginHandler.GET("/notes", notehandler.GetNote)
	ginHandler.GET("/notes/:id", notehandler.GetNotesByID)
	ginHandler.PUT("/notes/:id", notehandler.ModNote)
	ginHandler.DELETE("/notes/:id", notehandler.DelNote)

	ginHandler.POST("/register", userhandler.UserRegister)
	ginHandler.POST("/login", userhandler.UserLogin)

	ctx := context.Background()

	server := &http.Server{
		Addr:    ":8080",
		Handler: ginHandler,
	}

	if err := servers.RunServer(server, ctx, 10*time.Second); err != nil {
		log.Fatal(err)
	}
}
