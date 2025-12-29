package handlers

import (
	"net/http"

	"github.com/Yogi-1996/notes-backend/internal/models"
	"github.com/Yogi-1996/notes-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	service services.NoteServiceInterface
}

func NewNoteHandler(s services.NoteServiceInterface) *NoteHandler {
	return &NoteHandler{
		service: s,
	}
}

func (n *NoteHandler) NoteAdd(ctx *gin.Context) {
	var newnote models.Note

	if err := ctx.BindJSON(&newnote); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	note, err := n.service.AddNote(newnote.Title, newnote.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, note)
}

func (n *NoteHandler) GetNotes(ctx *gin.Context) {

	notes, err := n.service.GetAllNote()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, notes)
}
