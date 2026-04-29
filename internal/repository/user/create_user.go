package user

import (
	"context"

	"github.com/Khatchi/go-tweet/internal/model"
)

func (r *userRepository) CreateUser(ctx context.Context, model *model.UserModel) (int64, error) {
	query := `INSERT INTO users (email, username, password, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, model.Email, model.UserName, model.Password, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return 0, err
	}

	userID, _ := result.LastInsertId()

	return userID, nil
}
