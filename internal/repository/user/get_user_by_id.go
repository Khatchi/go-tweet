package user

import (
	"context"
	"database/sql"

	"github.com/Khatchi/go-tweet/internal/model"
)

func (r *userRepository) GetUSerByID(ctx context.Context, userID int64) (*model.UserModel, error) {
	query := `SELECT id, username, email, created_at, updated_at
		FROM users
		WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, userID)
	var result model.UserModel
	err := row.Scan(
		&result.ID,
		&result.UserName,
		&result.Email,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
