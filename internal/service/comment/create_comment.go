package comment

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Khatchi/go-tweet/internal/dto"
	"github.com/Khatchi/go-tweet/internal/model"
)

func (s *commentService) CreateComment(ctx context.Context, req *dto.StoreCommentRequest, userID int64) (int, error) {
	// check if tweet exists
	postExists, err := s.postRepo.GetPostByID(ctx, req.PostID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if postExists == nil {
		return http.StatusNotFound, errors.New("tweet not found")
	}

	// store comment
	now := time.Now()
	err = s.commentRepo.StoreComment(ctx, &model.CommentModel{
		PostID:    req.PostID,
		UserID:    userID,
		Content:   req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	// return
	return http.StatusCreated, nil
}
