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

func (repo *PostgresUserRepository) Create(ctx context.Context, user *entities.User) error {
	id := uuid.New()
	user.ID = id

	query := `INSERT INTO user_data (id, email, display_name) VALUES ($1, $2, $3)`

	_, err := repo.db.Exec(ctx, query, user.ID, user.Email, user.DisplayName)
	if err != nil {
		return fmt.Errorf("unable to insert row %w", err)
	}

	return nil
}

func (repo *PostgresUserRepository) FindAll(ctx context.Context) ([]*entities.User, error) {
	// TODO: implement pagination
	query := `
    SELECT 
        id,
        display_name,
        email,
        created_at,
        updated_at
    FROM
        user_data
    `

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	users := make([]*entities.User, 0)
	for rows.Next() {
		var user entities.User

		err := rows.Scan(
			&user.ID,
			&user.DisplayName,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		users = append(users, &user)
	}

	return users, nil
}

func (repo *PostgresUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return nil, nil
}

func (repo *PostgresUserRepository) Update(ctx context.Context, user *entities.User) error {
	return nil
}

func (repo *PostgresUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
