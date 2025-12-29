package services

import (
	"fmt"

	"github.com/Yogi-1996/notes-backend/internal/models"
	"github.com/Yogi-1996/notes-backend/internal/repository"
)

type NoteServiceInterface interface {
	AddNote(title, content string) (models.Note, error)
	GetNote(title string) (models.Note, error)
	GetAllNote() ([]models.Note, error)
}

type NoteService struct {
	repo *repository.NoteRepositry
}

func NewNoteService(n *repository.NoteRepositry) *NoteService {
	return &NoteService{
		repo: n,
	}
}

func (n *NoteService) AddNote(title, content string) (models.Note, error) {
	_, check := n.repo.GetNote(title)

	if !check {

		var newnote models.Note

		newnote = n.repo.Add(title, content)
		return newnote, nil

	}

	return models.Note{}, fmt.Errorf("Duplicate title")

}

func (n *NoteService) GetNote(title string) (models.Note, error) {
	note, check := n.repo.GetNote(title)

	if !check {

		return models.Note{}, fmt.Errorf("Duplicate title")

	}

	return note, nil

}

func (n *NoteService) GetAllNote() ([]models.Note, error) {
	notes, check := n.repo.GetAll()

	if !check {

		return []models.Note{}, fmt.Errorf("No Note Found")

	}

	return notes, nil
}
