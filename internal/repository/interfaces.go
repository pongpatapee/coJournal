package repository

import (
	"coJournal/internal/entities"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindAll(ctx context.Context) ([]*entities.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type JournalRepository interface {
	Create(ctx context.Context, journal *entities.Journal) error
	FindAll(ctx context.Context) ([]*entities.Journal, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Journal, error)
	Update(ctx context.Context, journal *entities.Journal) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type NoteRepository interface {
	Create(ctx context.Context, note *entities.Note) error
	FindAll(ctx context.Context) ([]*entities.Note, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Note, error)
	Update(ctx context.Context, note *entities.Note) error
	Delete(ctx context.Context, id uuid.UUID) error
}
