package repository

import (
	"github.com/Yogi-1996/notes-backend/internal/models"
	"gorm.io/gorm"
)

type NoteRepositryInterface interface {
	AddNote(note *models.Note) error
	ModNote(note_id, user_id int, note *models.Note) error
	DelNote(note_id, user_id int) error
	GetNote(note_id, user_id int) (models.Note, error)
	GetAllNote(user_id int) ([]models.Note, error)
}

type NoteRepositry struct {
	DB *gorm.DB
}

func NewNoteRepositry(db *gorm.DB) *NoteRepositry {
	return &NoteRepositry{
		DB: db,
	}
}

func (n *NoteRepositry) AddNote(note *models.Note) error {
	return n.DB.Create(note).Error
}

func (n *NoteRepositry) ModNote(note_id, user_id int, note *models.Note) error {
	var modnote models.Note

	if err := n.DB.Where("user_id = ? AND id = ?", user_id, note_id).First(&modnote).Error; err != nil {
		return err
	}

	n.DB.Model(&modnote).Updates(map[string]interface{}{
		"title":   note.Title,
		"content": note.Content,
	})

	return nil
}

func (n *NoteRepositry) DelNote(note_id, user_id int) error {
	var note models.Note

	err := n.DB.Where("user_id = ? AND id = ?", user_id, note_id).First(&note).Error
	if err != nil {
		return err
	}

	return n.DB.Delete(&note).Error
}

func (n *NoteRepositry) GetNote(note_id, user_id int) (models.Note, error) {
	var note models.Note

	err := n.DB.Where("user_id = ? AND id = ?", user_id, note_id).First(&note).Error

	return note, err
}

func (n *NoteRepositry) GetAllNote(user_id int) ([]models.Note, error) {
	var notes []models.Note

	err := n.DB.Where("user_id = ?", user_id).Find(&notes).Error

	return notes, err
}
