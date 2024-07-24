package service

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"

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

func (s *journalService) Create(journal *entities.Journal) error {
	journal.ID = uuid.New()
	return s.journalRepo.Create(journal)
}

func (s *journalService) FindAll() ([]*entities.Journal, error) {
	return s.journalRepo.FindAll()
}

func (s *journalService) FindByID(id uuid.UUID) (*entities.Journal, error) {
	return s.journalRepo.FindByID(id)
}

func (s *journalService) Update(journal *entities.Journal) error {
	return s.journalRepo.Update(journal)
}

func (s *journalService) Delete(id uuid.UUID) error {
	return s.journalRepo.Delete(id)
}
