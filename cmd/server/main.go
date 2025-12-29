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

	server.POST("/addnote", notehandler.NoteAdd)
	server.GET("/notes", notehandler.GetNotes)

	server.Run(":8080")

}
