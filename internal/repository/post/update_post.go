package post

import (
	"context"
	"errors"

	"github.com/Khatchi/go-tweet/internal/model"
)

func (r *postRepository) UpdatePost(ctx context.Context, model *model.PostModel, postID int64) error {
	query := `UPDATE posts SET title = ?, content = ?, updated_at = ?
		WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, model.Title, model.Content, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if rowAffected == 0 {
		return errors.New("no data to update")
	}

	return nil
}
