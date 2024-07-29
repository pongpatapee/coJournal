package service

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"
	"context"

	"github.com/google/uuid"
)

type journalService struct {
	journalRepo repository.JournalRepository
}

func NewJournalService(journalRepo repository.JournalRepository) JournalService {
	return &journalService{
		journalRepo: journalRepo,
	}
}

func (s *journalService) Create(ctx context.Context, journal *entities.Journal) error {
	return s.journalRepo.Create(ctx, journal)
}

func (s *journalService) FindAll(ctx context.Context) ([]*entities.Journal, error) {
	return s.journalRepo.FindAll(ctx)
}

func (s *journalService) FindByID(ctx context.Context, id uuid.UUID) (*entities.Journal, error) {
	return s.journalRepo.FindByID(ctx, id)
}

func (s *journalService) Update(ctx context.Context, journal *entities.Journal) error {
	return s.journalRepo.Update(ctx, journal)
}

func (s *journalService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.journalRepo.Delete(ctx, id)
}
