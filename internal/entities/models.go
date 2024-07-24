package entities

import (
	"github.com/google/uuid"
)

type User struct {
	// CreateDate  time.Time `json:"create_date"`
	// UpdatedDate time.Time `json:"updated_date"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	ID          uuid.UUID `json:"id"`
}
