package store

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
    INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id, created_at
  `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) DeleteSeedAll(ctx context.Context) error {
	query := `DELETE FROM users`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
			return err
	}

	// rows, err := result.RowsAffected()
	// if err != nil {
	// 		return err
	// }

	// if rows == 0 {
	// 		return ErrNotFound
	// }
	fmt.Println("Successfully cleaned comments table!")
	return nil
}
