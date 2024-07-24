package in_memory

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"
	"errors"
	"sync"

	"github.com/google/uuid"
)

type InMemoryUserRepository struct {
	db map[uuid.UUID]*entities.User
	sync.Mutex
}

func NewInMemoryUserRepository() repository.UserRepository {
	return &InMemoryUserRepository{
		db: make(map[uuid.UUID]*entities.User),
	}
}

func (repo *InMemoryUserRepository) Create(user *entities.User) error {
	repo.Lock()
	defer repo.Unlock()

	id := uuid.New()
	user.ID = id
	repo.db[id] = user

	return nil
}

func (repo *InMemoryUserRepository) FindAll() ([]*entities.User, error) {
	repo.Lock()
	defer repo.Unlock()

	users := make([]*entities.User, 0, len(repo.db))

	for _, userPtr := range repo.db {
		users = append(users, userPtr)
	}

	return users, nil
}

func (repo *InMemoryUserRepository) FindByID(id uuid.UUID) (*entities.User, error) {
	repo.Lock()
	defer repo.Unlock()

	user, found := repo.db[id]

	if !found {
		return nil, errors.New("id does not exist")
	}

	return user, nil
}

func (repo *InMemoryUserRepository) Update(user *entities.User) error {
	repo.Lock()
	defer repo.Unlock()

	id := user.ID

	_, found := repo.db[id]
	if !found {
		return errors.New("id does not exist")
	}

	repo.db[id] = user

	return nil
}

func (repo *InMemoryUserRepository) Delete(id uuid.UUID) error {
	repo.Lock()
	defer repo.Unlock()

	_, found := repo.db[id]
	if !found {
		return errors.New("id does not exist")
	}

	delete(repo.db, id)

	return nil
}
