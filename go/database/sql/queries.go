package sql

import (
	"context"
	"database/sql"
	"time"
)

// User is sample structure.
type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Repo is a repository for a DB.
type Repo struct {
	db *sql.DB
}

// NewRepo creates a new Repo.
func NewRepo(db *sql.DB) *Repo {
	return &Repo{db}
}

// FindUser gets user from repository.
func (repo *Repo) FindUser(ctx context.Context, id int64) (*User, error) {
	u := &User{}
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = conn.QueryRowContext(ctx, `
		SELECT id, name, email, created_at, updated_at FROM user WHERE id = ?
	`, id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return u, nil
}

// AddUser adds user to repository.
func (repo *Repo) AddUser(ctx context.Context, u *User) error {
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	now := time.Now()
	res, err := conn.ExecContext(ctx, `
        INSERT INTO user (name, email, created_at, updated_at)
        VALUES (?, ?, ?, ?)
    `, u.Name, u.Email, now, now,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId() // 挿入した行のIDを返却
	if err != nil {
		return err
	}
	u.ID = id
	u.CreatedAt = now
	u.UpdatedAt = now

	return nil
}
