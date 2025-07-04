package user

import (
	"context"
	"database/sql"
)

type Repository interface {
	Save(ctx context.Context, user User) (User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(ctx context.Context, user User) (User, error) {
	query := `INSERT INTO users (name, occupation, email, password_hash, role, created_at, updated_at)
				VALUES($1, $2, $3, $4, $5, NOW(), NOW())`

	_, err := r.db.ExecContext(ctx, query, user.Name, user.Occupation, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		return user, err
	}

	return user, nil
}
