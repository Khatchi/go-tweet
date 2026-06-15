package comment

import (
	"context"

	"github.com/Khatchi/go-tweet/internal/config"
	"github.com/Khatchi/go-tweet/internal/dto"
	"github.com/Khatchi/go-tweet/internal/repository/comment"
	"github.com/Khatchi/go-tweet/internal/repository/post"
)

type CommentService interface {
	CreateComment(ctx context.Context, req *dto.StoreCommentRequest, userID int64) (int, error)
}

type commentService struct {
	cfg         *config.Config
	commentRepo comment.CommentRepository
	postRepo    post.PostRepository
}

func NewCommentService(cfg *config.Config, commentRepo comment.CommentRepository, postRepo post.PostRepository) CommentService {
	return &commentService{
		cfg:         cfg,
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}
