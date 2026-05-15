package post

import (
	"context"
	"database/sql"
	"time"

	"github.com/Khatchi/go-tweet/internal/model"
)

type PostRepository interface {
	StorePost(ctx context.Context, model *model.PostModel) (int64, error)
	GetPostByID(ctx context.Context, postID int64) (*model.PostModel, error)
	UpdatePost(ctx context.Context, model *model.PostModel, postID int64) error
	SoftDeletePost(ctx context.Context, postID int64, now time.Time) error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}
