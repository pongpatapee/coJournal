package in_memory

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"
	"errors"
	"sync"

	"github.com/google/uuid"
)

type InMemoryNoteRepository struct {
	db map[uuid.UUID]*entities.Note
	sync.Mutex
}

func NewInMemoryNoteRepository() repository.NoteRepository {
	return &InMemoryNoteRepository{
		db: make(map[uuid.UUID]*entities.Note),
	}
}

func (repo *InMemoryNoteRepository) Create(note *entities.Note) error {
	repo.Lock()
	defer repo.Unlock()

	id := uuid.New()
	note.ID = id
	repo.db[id] = note

	return nil
}

func (repo *InMemoryNoteRepository) FindAll() ([]*entities.Note, error) {
	repo.Lock()
	defer repo.Unlock()

	notes := make([]*entities.Note, 0, len(repo.db))

	for _, notePtr := range repo.db {
		notes = append(notes, notePtr)
	}

	return notes, nil
}

func (repo *InMemoryNoteRepository) FindByID(id uuid.UUID) (*entities.Note, error) {
	repo.Lock()
	defer repo.Unlock()

	note, found := repo.db[id]

	if !found {
		return nil, errors.New("id does not exist")
	}

	return note, nil
}

func (repo *InMemoryNoteRepository) Update(note *entities.Note) error {
	repo.Lock()
	defer repo.Unlock()

	id := note.ID

	_, found := repo.db[id]
	if !found {
		return errors.New("id does not exist")
	}

	repo.db[id] = note

	return nil
}

func (repo *InMemoryNoteRepository) Delete(id uuid.UUID) error {
	repo.Lock()
	defer repo.Unlock()

	_, found := repo.db[id]
	if !found {
		return errors.New("id does not exist")
	}

	delete(repo.db, id)

	return nil
}
