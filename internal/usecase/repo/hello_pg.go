package repo

import (
	"charitybay/internal/entity"
	"charitybay/pkg/db"
	"context"
	"fmt"
)

type GreetRepo struct {
	*db.Postgres
}

// NewGreetRepo creates a new NewGreetRepo
func New(db *db.Postgres) *GreetRepo {
	return &GreetRepo{db}
}

// GetGreets fetches all greets
func (r *GreetRepo) GetGreets(ctx context.Context) ([]entity.Greet, error) {
	sql, _, err := r.Builder.Select("message").From("greetings").ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetGreets - ToSql: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("GetGreets - Query: %w", err)
	}
	defer rows.Close()

	entites := []entity.Greet{}

	for rows.Next() {
		var e entity.Greet
		err = rows.Scan(&e.Message)
		if err != nil {
			return nil, fmt.Errorf("GetGreets - Scan: %w", err)
		}

		entites = append(entites, e)
	}

	return entites, nil
}

// CreateGreet creates a new greet
func (r *GreetRepo) CreateGreet(ctx context.Context, greet entity.Greet) error {
	sql, args, err := r.Builder.Insert("greetings").Columns("message").Values(greet.Message).ToSql()
	if err != nil {
		return fmt.Errorf("CreateGreet - ToSql: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return fmt.Errorf("CreateGreet - Exec: %w", err)
	}

	return nil
}
