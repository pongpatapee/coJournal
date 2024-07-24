package service

import (
	"coJournal/internal/entities"

	"github.com/google/uuid"
)

type UserService interface {
	Create(user *entities.User) error
	FindAll() ([]*entities.User, error)
	FindByID(id uuid.UUID) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uuid.UUID) error
}

type JournalService interface {
	Create(journal *entities.Journal) error
	FindAll() ([]*entities.Journal, error)
	FindByID(id uuid.UUID) (*entities.Journal, error)
	Update(journal *entities.Journal) error
	Delete(id uuid.UUID) error
}

type NoteService interface {
	Create(note *entities.Note) error
	FindAll() ([]*entities.Note, error)
	FindByID(id uuid.UUID) (*entities.Note, error)
	Update(note *entities.Note) error
	Delete(id uuid.UUID) error
}

// type Service struct {
// 	UserService    UserService
// 	JournalService JournalService
// 	NoteService    NoteService
// }
