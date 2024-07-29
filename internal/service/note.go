package service

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"
	"context"

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

func (s *noteService) Create(ctx context.Context, note *entities.Note) error {
	return s.noteRepo.Create(ctx, note)
}

func (s *noteService) FindAll(ctx context.Context) ([]*entities.Note, error) {
	return s.noteRepo.FindAll(ctx)
}

func (s *noteService) FindByID(ctx context.Context, id uuid.UUID) (*entities.Note, error) {
	return s.noteRepo.FindByID(ctx, id)
}

func (s *noteService) Update(ctx context.Context, note *entities.Note) error {
	return s.noteRepo.Update(ctx, note)
}

func (s *noteService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.noteRepo.Delete(ctx, id)
}
