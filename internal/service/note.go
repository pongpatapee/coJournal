package service

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"

	"github.com/google/uuid"
)

type noteService struct {
	noteRepo repository.NoteRepository
}

func NewNoteService(noteRepo repository.NoteRepository) NoteService {
	return &noteService{
		noteRepo: noteRepo,
	}
}

func (s *noteService) Create(note *entities.Note) error {
	note.ID = uuid.New()
	return s.noteRepo.Create(note)
}

func (s *noteService) FindAll() ([]*entities.Note, error) {
	return s.noteRepo.FindAll()
}

func (s *noteService) FindByID(id uuid.UUID) (*entities.Note, error) {
	return s.noteRepo.FindByID(id)
}

func (s *noteService) Update(note *entities.Note) error {
	return s.noteRepo.Update(note)
}

func (s *noteService) Delete(id uuid.UUID) error {
	return s.noteRepo.Delete(id)
}
