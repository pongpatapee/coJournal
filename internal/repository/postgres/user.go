package postgres

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

	query := `INSERT INTO user_data (id, email, display_name) VALUES (@id, @email, @display_name)`
	args := pgx.NamedArgs{
		"id":           user.ID,
		"email":        user.Email,
		"display_name": user.DisplayName,
	}

	_, err := repo.db.Exec(ctx, query, args)
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
	query := `
    SELECT 
        id,
        display_name,
        email,
        created_at,
        updated_at
    FROM
        user_data
    WHERE
        id = @id
    `
	args := pgx.NamedArgs{
		"id": id,
	}

	var user entities.User
	err := repo.db.QueryRow(ctx, query, args).Scan(
		&user.ID,
		&user.DisplayName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *PostgresUserRepository) Update(ctx context.Context, user *entities.User) error {
	query := `
    UPDATE 
        user_data 
    SET 
        email=@email,
        display_name=@display_name,
        updated_at=NOW() 
    WHERE 
        id=@id
    `
	args := pgx.NamedArgs{
		"id":           user.ID,
		"email":        user.Email,
		"display_name": user.DisplayName,
	}
	_, err := repo.db.Exec(ctx, query, args)

	return err
}

func (repo *PostgresUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM user_data WHERE id=@id`
	_, err := repo.db.Exec(ctx, query, pgx.NamedArgs{"id": id})
	return err
}
