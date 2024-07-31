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

type PostgresNoteRepository struct {
	db *pgxpool.Pool
}

func NewPostgresNoteRepository(db *pgxpool.Pool) repository.NoteRepository {
	return &PostgresNoteRepository{
		db: db,
	}
}

func (repo *PostgresNoteRepository) Create(ctx context.Context, note *entities.Note) error {
	id := uuid.New()
	note.ID = id

	query := `
    INSERT INTO
        note (id, journal_id, author, title, body, last_viewed) 
    VALUES
        (@id, @journal_id, @author, @title, @body, @last_viewed)
    `

	args := pgx.NamedArgs{
		"id":          note.ID,
		"journal_id":  note.JournalID,
		"author":      note.Author,
		"title":       note.Title,
		"body":        note.Body,
		"last_viewed": nil,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row %w", err)
	}

	return nil
}

func (repo *PostgresNoteRepository) FindAll(ctx context.Context) ([]*entities.Note, error) {
	// TODO: implement pagination
	query := `
    SELECT 
        id,
        journal_id,
        author,
        title,
        body,
        last_viewed,
        updated_at,
        created_at
    FROM
        note
    `

	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	notes := make([]*entities.Note, 0)
	for rows.Next() {
		var note entities.Note

		err := rows.Scan(
			&note.ID,
			&note.JournalID,
			&note.Author,
			&note.Title,
			&note.Body,
			&note.LastViewed,
			&note.UpdatedAt,
			&note.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		notes = append(notes, &note)
	}

	return notes, nil
}

func (repo *PostgresNoteRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Note, error) {
	query := `
    SELECT 
        id,
        journal_id,
        author,
        title,
        body,
        last_viewed,
        updated_at,
        created_at
    FROM
        note
    WHERE
        id = @id
    `
	args := pgx.NamedArgs{
		"id": id,
	}

	var note entities.Note
	err := repo.db.QueryRow(ctx, query, args).Scan(
		&note.ID,
		&note.JournalID,
		&note.Author,
		&note.Title,
		&note.Body,
		&note.LastViewed,
		&note.UpdatedAt,
		&note.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (repo *PostgresNoteRepository) Update(ctx context.Context, note *entities.Note) error {
	query := `
    UPDATE 
        note 
    SET 
        title=@title,
        body=@body,
        last_viewed=@last_viewed,
        updated_at=NOW()

    WHERE 
        id=@id
    `
	args := pgx.NamedArgs{
		"id":          note.ID,
		"title":       note.Title,
		"body":        note.Body,
		"last_viewed": note.LastViewed,
	}
	_, err := repo.db.Exec(ctx, query, args)

	return err
}

func (repo *PostgresNoteRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM note WHERE id=@id`
	_, err := repo.db.Exec(ctx, query, pgx.NamedArgs{"id": id})
	return err
}
