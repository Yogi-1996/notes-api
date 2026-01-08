package repository

import (
	"sync"
	"time"

	"github.com/Yogi-1996/notes-backend/internal/models"
)

type NoteRepositryInterface interface {
	AddNote(user_id int, title, content string) (models.Note, bool)
	ModNote(user_id, id int, note models.Note) (models.Note, bool)
	DelNote(user_id, id int) bool
	GetNote(user_id, id int) (models.Note, bool)
	GetAll(user_id int) ([]models.Note, bool)
}

type NoteRepositry struct {
	mu     sync.Mutex
	notes  map[int]models.Note
	nextID int
}

func NewNoteRepositry() *NoteRepositry {
	return &NoteRepositry{
		notes:  make(map[int]models.Note),
		nextID: 1,
	}
}

func (n *NoteRepositry) AddNote(user_id int, title, content string) (models.Note, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, note := range n.notes {
		if note.Title == title {
			return models.Note{}, false
		}
	}

	newnote := models.Note{
		ID:        n.nextID,
		UserID:    user_id,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	n.notes[n.nextID] = newnote
	n.nextID += 1

	return newnote, true
}

func (n *NoteRepositry) ModNote(user_id, id int, note models.Note) (models.Note, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	existing, ok := n.notes[id]
	if !ok {
		return models.Note{}, false
	}

	existing.Content = note.Content
	existing.Title = note.Title
	existing.UpdatedAt = time.Now()

	n.notes[id] = existing

	return existing, true
}

func (n *NoteRepositry) DelNote(user_id, id int) bool {
	n.mu.Lock()
	defer n.mu.Unlock()

	if note, ok := n.notes[id]; !ok {
		return false
	} else if note.UserID != user_id {
		return false
	}

	delete(n.notes, id)
	return true
}

func (n *NoteRepositry) GetNote(user_id, id int) (models.Note, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, note := range n.notes {
		if note.ID == id && note.UserID == user_id {
			return note, true
		}
	}

	return models.Note{}, false
}

func (n *NoteRepositry) GetAll(user_id int) ([]models.Note, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	notes := make([]models.Note, 0, len(n.notes))
	for _, note := range n.notes {
		if note.UserID == user_id {
			notes = append(notes, note)
		}

	}

	if len(notes) == 0 {
		return []models.Note{}, false
	}

	return notes, true
}
