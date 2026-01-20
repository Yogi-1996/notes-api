package services

import (
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
	newnote := models.Note{
		Title:   title,
		Content: content,
		UserID:  user_id,
	}

	err := n.repo.AddNote(&newnote)
	if err != nil {
		return models.Note{}, err
	}

	return newnote, nil
}

func (n *NoteService) ModNote(user_id, id int, note models.Note) (models.Note, error) {
	err := n.repo.ModNote(id, user_id, &note)
	if err != nil {
		return models.Note{}, err
	}

	return note, nil
}

func (n *NoteService) DelNote(user_id, id int) error {
	err := n.repo.DelNote(id, user_id)
	if err != nil {
		return err
	}

	return nil

}

func (n *NoteService) GetNote(user_id, id int) (models.Note, error) {
	note, err := n.repo.GetNote(id, user_id)
	if err != nil {
		return models.Note{}, err
	}

	return note, nil

}

func (n *NoteService) GetAll(user_id int) ([]models.Note, error) {
	notes, err := n.repo.GetAllNote(user_id)
	if err != nil {
		return []models.Note{}, err
	}

	return notes, nil

}
