package repository

import (
	"sync"
	"time"

	"github.com/Yogi-1996/notes-backend/internal/models"
)

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

func (n *NoteRepositry) Add(title, content string) models.Note {
	n.mu.Lock()
	defer n.mu.Unlock()

	newnote := models.Note{
		ID:        n.nextID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	n.notes[n.nextID] = newnote
	n.nextID += 1

	return newnote
}

func (n *NoteRepositry) GetNote(title string) (models.Note, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, note := range n.notes {
		if note.Title == title {
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
