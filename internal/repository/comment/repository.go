package comment

import (
	"context"
	"database/sql"

	"github.com/Khatchi/go-tweet/internal/model"
)

type CommentRepository interface {
	StoreComment(ctx context.Context, model *model.CommentModel) error
}

type commentRepository struct {
	db *sql.DB
}

func NewRepostory(db *sql.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}
