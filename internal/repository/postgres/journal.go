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
	return err
}

func (repo *PostgresJournalRepository) FindAll(ctx context.Context) ([]*entities.Journal, error) {
	query := `
    SELECT
        id,
        name,
        user_a,
        user_b,
        created_at,
        updated_at
    FROM
        journal
    `

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	journals := make([]*entities.Journal, 0)
	for rows.Next() {
		var journal entities.Journal

		rows.Scan(
			&journal.ID,
			&journal.Name,
			&journal.UserA,
			&journal.UserB,
			&journal.CreatedAt,
			&journal.UpdatedAt,
		)

		journals = append(journals, &journal)
	}

	return journals, nil
}

func (repo *PostgresJournalRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Journal, error) {
	query := `
    SELECT
        id,
        name,
        user_a,
        user_b,
        created_at,
        updated_at
    FROM
        journal
    WHERE
        id=@id
    `
	var journal entities.Journal
	err := repo.db.QueryRow(ctx, query, pgx.NamedArgs{"id": id}).
		Scan(
			&journal.ID,
			&journal.Name,
			&journal.UserA,
			&journal.UserB,
			&journal.CreatedAt,
			&journal.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}

	return &journal, nil
}

func (repo *PostgresJournalRepository) Update(ctx context.Context, journal *entities.Journal) error {
	query := `
    UPDATE 
        journal
    SET
        name=@name,
        user_a=@user_a,
        user_b=@user_b,
        updated_at=NOW()
    WHERE
        id=@id
    `

	args := pgx.NamedArgs{
		"id":     journal.ID,
		"name":   journal.Name,
		"user_a": journal.UserA,
		"user_b": journal.UserB,
	}

	_, err := repo.db.Exec(ctx, query, args)

	return err
}

func (repo *PostgresJournalRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
    DELETE FROM 
        journal 
    WHERE 
        id=@id
    `
	_, err := repo.db.Exec(ctx, query, pgx.NamedArgs{"id": id})

	return err
}
