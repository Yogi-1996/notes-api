package services

import (
	"fmt"

	"github.com/Yogi-1996/notes-backend/internal/models"
	"github.com/Yogi-1996/notes-backend/internal/repository"
)

type NoteServiceInterface interface {
	AddNote(user_id int, title, content string) (models.Note, error)
	ModNote(user_id, id int, note models.Note) (models.Note, error)
	DelNote(user_id, id int) error
	GetNote(user_id, id int) (models.Note, error)
	GetAll(user_id int) ([]models.Note, error)
}

type NoteService struct {
	repo repository.NoteRepositryInterface
}

func NewNoteService(n repository.NoteRepositryInterface) *NoteService {
	return &NoteService{
		repo: n,
	}
}

func (n *NoteService) AddNote(user_id int, title, content string) (models.Note, error) {
	note, check := n.repo.AddNote(user_id, title, content)
	if !check {
		return models.Note{}, fmt.Errorf("Duplicate Note Title")
	}

	return note, nil
}

func (n *NoteService) ModNote(user_id, id int, note models.Note) (models.Note, error) {
	note, check := n.repo.ModNote(user_id, id, note)
	if !check {
		return models.Note{}, fmt.Errorf("Id Not Found")
	}

	return note, nil
}

func (n *NoteService) DelNote(user_id, id int) error {
	check := n.repo.DelNote(user_id, id)
	if !check {
		return fmt.Errorf("Id Not Found")
	}

	return nil

}

func (n *NoteService) GetNote(user_id, id int) (models.Note, error) {
	note, check := n.repo.GetNote(user_id, id)
	if !check {
		return models.Note{}, fmt.Errorf("Id Not Found")
	}

	return note, nil

}

func (n *NoteService) GetAll(user_id int) ([]models.Note, error) {
	note, check := n.repo.GetAll(user_id)
	if !check {
		return []models.Note{}, fmt.Errorf("Notes Empty")
	}

	return note, nil

}
