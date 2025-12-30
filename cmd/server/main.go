package main

import (
	"net/http"

	"github.com/Yogi-1996/notes-backend/internal/handlers"
	"github.com/Yogi-1996/notes-backend/internal/repository"
	"github.com/Yogi-1996/notes-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {

	notesdb := repository.NewNoteRepositry()
	noteservice := services.NewNoteService(notesdb)
	notehandler := handlers.NewNoteHandler(noteservice)

	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"messsage": "Note Server Running",
		})
	})

	server.POST("/notes", notehandler.NoteAdd)
	server.GET("/notes", notehandler.GetNote)
	server.GET("/notes/:id", notehandler.GetNotesByID)
	server.PUT("/notes/:id", notehandler.ModNote)
	server.DELETE("/notes/:id", notehandler.DelNote)

	server.Run(":8080")

}
