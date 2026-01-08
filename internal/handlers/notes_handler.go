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

func (n *NoteHandler) NoteAdd(c *gin.Context) {
	var newnote models.Note

	userIDValue, ok := c.Get("UserID")

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "User not Auntheticated",
		})
		return
	}

	user_id, ok := userIDValue.(int)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "userID has wrong type",
		})
		return
	}

	if err := c.ShouldBindJSON(&newnote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	note, err := n.service.AddNote(user_id, newnote.Title, newnote.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, note)
}

func (n *NoteHandler) ModNote(c *gin.Context) {
	notestr := c.Param("id")

	userIDValue, ok := c.Get("UserID")

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "User not Auntheticated",
		})
		return
	}

	user_id, ok := userIDValue.(int)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "userID has wrong type",
		})
		return
	}

	noteid, err := strconv.Atoi(notestr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid note id",
			"error":   err.Error(),
		})
		return
	}

	var newnote models.Note

	if err := c.ShouldBindJSON(&newnote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	note, err := n.service.ModNote(user_id, noteid, newnote)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Data":    note,
		"message": "note Modified Sucessfully",
	})

}

func (n *NoteHandler) DelNote(c *gin.Context) {

	notestr := c.Param("id")

	userIDValue, ok := c.Get("UserID")

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "User not Auntheticated",
		})
		return
	}

	user_id, ok := userIDValue.(int)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "userID has wrong type",
		})
		return
	}

	noteid, err := strconv.Atoi(notestr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid note id",
			"error":   err.Error(),
		})
		return
	}

	err = n.service.DelNote(user_id, noteid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Deleted Sucessfully",
	})

}

func (n *NoteHandler) GetNotesByID(c *gin.Context) {

	notestr := c.Param("id")

	userIDValue, ok := c.Get("UserID")

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "User not Auntheticated",
		})
		return
	}

	user_id, ok := userIDValue.(int)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "userID has wrong type",
		})
		return
	}

	noteid, err := strconv.Atoi(notestr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid note id",
			"error":   err.Error(),
		})
		return
	}

	note, err := n.service.GetNote(user_id, noteid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Data":    note,
		"message": "note retrieved Sucessfully",
	})

}

func (n *NoteHandler) GetNote(c *gin.Context) {
	userIDValue, ok := c.Get("UserID")

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "User not Auntheticated",
		})
		return
	}

	user_id, ok := userIDValue.(int)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "userID has wrong type",
		})
		return
	}

	notes, err := n.service.GetAll(user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Data":    notes,
		"message": "note retrieved Sucessfully",
	})
}
