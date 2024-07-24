package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreateDate  time.Time `json:"create_date"`
	UpdatedDate time.Time `json:"updated_date"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	ID          uuid.UUID `json:"user_id"`
}

type Journal struct {
	CreateDate  time.Time `json:"create_date"`
	UpdatedDate time.Time `json:"updated_date"`
	Name        string    `json:"name"`
	Notes       []Note    `json:"notes"`
	UserA       uuid.UUID `json:"user_a"`
	UserB       uuid.UUID `json:"user_b"`
	ID          uuid.UUID `json:"journal_id"`
}

type Note struct {
	CreateDate  time.Time `json:"create_date"`
	UpdatedDate time.Time `json:"updated_date"`
	LastViewed  time.Time `json:"last_viewed"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Author      uuid.UUID `json:"author"`
	ID          uuid.UUID `json:"note_id"`
}
