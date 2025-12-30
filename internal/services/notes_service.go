package services

import (
	"fmt"

	"github.com/Yogi-1996/notes-backend/internal/models"
	"github.com/Yogi-1996/notes-backend/internal/repository"
)

type NoteServiceInterface interface {
	AddNote(title, content string) (models.Note, error)
	ModNote(id int, note models.Note) (models.Note, error)
	DelNote(id int) error
	GetNote(id int) (models.Note, error)
	GetAll() ([]models.Note, error)
}

type NoteService struct {
	repo repository.NoteRepositryInterface
}

func NewNoteService(n repository.NoteRepositryInterface) *NoteService {
	return &NoteService{
		repo: n,
	}
}

func (n *NoteService) AddNote(title, content string) (models.Note, error) {
	note, check := n.repo.AddNote(title, content)
	if !check {
		return models.Note{}, fmt.Errorf("Duplicate Note Title")
	}

	return note, nil
}

func (n *NoteService) ModNote(id int, note models.Note) (models.Note, error) {
	note, check := n.repo.ModNote(id, note)
	if !check {
		return models.Note{}, fmt.Errorf("Id Not Found")
	}

	return note, nil
}

func (n *NoteService) DelNote(id int) error {
	check := n.repo.DelNote(id)
	if !check {
		return fmt.Errorf("Id Not Found")
	}

	return nil

}

func (n *NoteService) GetNote(id int) (models.Note, error) {
	note, check := n.repo.GetNote(id)
	if !check {
		return models.Note{}, fmt.Errorf("Id Not Found")
	}

	return note, nil

}

func (n *NoteService) GetAll() ([]models.Note, error) {
	note, check := n.repo.GetAll()
	if !check {
		return []models.Note{}, fmt.Errorf("Notes Empty")
	}

	return note, nil

}
