package postgres

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) repository.UserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (repo *PostgresUserRepository) Create(user *entities.User) error {
	id := uuid.New()
	user.ID = id

	query := `INSERT INTO user (id, email, displayname)
        VALUES ($1, $2, $3)
        `

	_, err := repo.db.Exec(context.Background(), query, user.ID, user.Email, user.DisplayName)
	if err != nil {
		return fmt.Errorf("unable to insert row %w", err)
	}

	return nil
}

func (repo *PostgresUserRepository) FindAll() ([]*entities.User, error) {
	return []*entities.User{}, nil
}

func (repo *PostgresUserRepository) FindByID(id uuid.UUID) (*entities.User, error) {
	return nil, nil
}

func (repo *PostgresUserRepository) Update(user *entities.User) error {
	return nil
}

func (repo *PostgresUserRepository) Delete(id uuid.UUID) error {
	return nil
}
