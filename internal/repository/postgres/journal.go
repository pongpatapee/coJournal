package postgres

import (
	"coJournal/internal/entities"
	"coJournal/internal/repository"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresJournalRepository struct {
	db *pgxpool.Pool
}

func NewPostgresJournalRepository(db *pgxpool.Pool) repository.JournalRepository {
	return &PostgresJournalRepository{
		db: db,
	}
}

func (repo *PostgresJournalRepository) Create(ctx context.Context, journal *entities.Journal) error {
	journal.ID = uuid.New()
	query := `INSERT INTO journal (id, name, user_a, user_b) VALUES (@id, @name, @user_a, @user_b)`
	args := pgx.NamedArgs{
		"id":     journal.ID,
		"name":   journal.Name,
		"user_a": journal.UserA,
		"user_b": journal.UserB,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}

func (repo *PostgresJournalRepository) FindAll(ctx context.Context) ([]*entities.Journal, error) {
	journals := make([]*entities.Journal, 0)

	return []*entities.Journal{}, nil
}

func (repo *PostgresJournalRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Journal, error) {
	return nil, nil
}

func (repo *PostgresJournalRepository) Update(ctx context.Context, journal *entities.Journal) error {
	return nil
}

func (repo *PostgresJournalRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
