package in_memory

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"
	"errors"
	"sync"

	"github.com/google/uuid"
)

type InMemoryJournalRepository struct {
	db map[uuid.UUID]*entities.Journal
	sync.Mutex
}

func NewInMemoryJournalRepository() repository.JournalRepository {
	return &InMemoryJournalRepository{
		db: make(map[uuid.UUID]*entities.Journal),
	}
}

func (repo *InMemoryJournalRepository) Create(journal *entities.Journal) error {
	repo.Lock()
	defer repo.Unlock()

	id := uuid.New()
	journal.ID = id
	repo.db[id] = journal

	return nil
}

func (repo *InMemoryJournalRepository) FindAll() ([]*entities.Journal, error) {
	repo.Lock()
	defer repo.Unlock()

	journals := make([]*entities.Journal, 0, len(repo.db))

	for _, journalPtr := range repo.db {
		journals = append(journals, journalPtr)
	}

	return journals, nil
}

func (repo *InMemoryJournalRepository) FindByID(id uuid.UUID) (*entities.Journal, error) {
	repo.Lock()
	defer repo.Unlock()

	journal, found := repo.db[id]

	if !found {
		return nil, errors.New("id does not exist")
	}

	return journal, nil
}

func (repo *InMemoryJournalRepository) Update(journal *entities.Journal) error {
	repo.Lock()
	defer repo.Unlock()

	id := journal.ID

	_, found := repo.db[id]
	if !found {
		return errors.New("id does not exist")
	}

	repo.db[id] = journal

	return nil
}

func (repo *InMemoryJournalRepository) Delete(id uuid.UUID) error {
	repo.Lock()
	defer repo.Unlock()

	_, found := repo.db[id]
	if !found {
		return errors.New("id does not exist")
	}

	delete(repo.db, id)

	return nil
}
