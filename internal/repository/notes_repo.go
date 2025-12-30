package repository

import (
	"sync"
	"time"

	"github.com/Yogi-1996/notes-backend/internal/models"
)

type NoteRepositryInterface interface {
	AddNote(title, content string) (models.Note, bool)
	ModNote(id int, note models.Note) (models.Note, bool)
	DelNote(id int) bool
	GetNote(id int) (models.Note, bool)
	GetAll() ([]models.Note, bool)
}

type NoteRepositry struct {
	mu     sync.Mutex
	notes  map[int]models.Note
	nextID int
}

func NewNoteRepositry() *NoteRepositry {
	return &NoteRepositry{
		notes: make(map[int]models.Note),
	}
}

func (n *NoteRepositry) AddNote(title, content string) (models.Note, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, note := range n.notes {
		if note.Title == title {
			return models.Note{}, false
		}
	}

	newnote := models.Note{
		ID:        n.nextID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	n.notes[n.nextID] = newnote
	n.nextID += 1

	return newnote, true
}

func (n *NoteRepositry) ModNote(id int, note models.Note) (models.Note, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if _, ok := n.notes[id]; !ok {
		return models.Note{}, false
	}

	note.ID = id
	n.notes[id] = note

	return note, true
}

func (n *NoteRepositry) DelNote(id int) bool {
	n.mu.Lock()
	defer n.mu.Unlock()

	if _, ok := n.notes[id]; !ok {
		return false
	}

	delete(n.notes, id)
	return true
}

func (n *NoteRepositry) GetNote(id int) (models.Note, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, note := range n.notes {
		if note.ID == id {
			return note, true
		}
	}

	return models.Note{}, false
}

func (n *NoteRepositry) GetAll() ([]models.Note, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if len(n.notes) == 0 {
		return []models.Note{}, false
	}

	notes := make([]models.Note, 0, len(n.notes))
	for _, note := range n.notes {
		notes = append(notes, note)
	}
	return notes, true
}
