package handlers

import (
	"net/http"
	"strconv"

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

func (n *NoteHandler) ModNote(ctx *gin.Context) {
	notestr := ctx.Param("id")

	noteid, err := strconv.Atoi(notestr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid note id",
			"error":   err.Error(),
		})
		return
	}

	var newnote models.Note

	if err := ctx.BindJSON(&newnote); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	note, err := n.service.ModNote(noteid, newnote)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Data":    note,
		"message": "note Modified Sucessfully",
	})

}

func (n *NoteHandler) DelNote(ctx *gin.Context) {

	notestr := ctx.Param("id")

	noteid, err := strconv.Atoi(notestr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid note id",
			"error":   err.Error(),
		})
		return
	}

	err = n.service.DelNote(noteid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Deleted Sucessfully",
	})

}

func (n *NoteHandler) GetNotesByID(ctx *gin.Context) {

	notestr := ctx.Param("id")

	noteid, err := strconv.Atoi(notestr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid note id",
			"error":   err.Error(),
		})
		return
	}

	note, err := n.service.GetNote(noteid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Data":    note,
		"message": "note retrieved Sucessfully",
	})

}

func (n *NoteHandler) GetNote(ctx *gin.Context) {
	notes, err := n.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Data":    notes,
		"message": "note retrieved Sucessfully",
	})
}
