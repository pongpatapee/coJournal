package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	ID          uuid.UUID `json:"user_id"`
}

type Journal struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Notes     []Note    `json:"notes"`
	UserA     uuid.UUID `json:"user_a"`
	UserB     uuid.UUID `json:"user_b"`
	ID        uuid.UUID `json:"journal_id"`
}

type Note struct {
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastViewed time.Time `json:"last_viewed"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Author     uuid.UUID `json:"author"`
	JournalID  uuid.UUID `json:"journal_id"`
	ID         uuid.UUID `json:"note_id"`
}
