package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DisplayName string    `json:"display_name" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	ID          uuid.UUID `json:"user_id"`
}

type Journal struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" validate:"required"`
	Notes     []Note    `json:"notes"`
	UserA     uuid.UUID `json:"user_a" validate:"required"`
	UserB     uuid.UUID `json:"user_b"`
	ID        uuid.UUID `json:"journal_id"`
}

type Note struct {
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastViewed time.Time `json:"last_viewed"`
	Title      string    `json:"title" validate:"required"`
	Body       string    `json:"body"`
	Author     uuid.UUID `json:"author" validate:"required"`
	JournalID  uuid.UUID `json:"journal_id" validate:"required"`
	ID         uuid.UUID `json:"note_id"`
}
