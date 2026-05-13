package post

import (
	"context"
	"net/http"
	"time"

	"github.com/Khatchi/go-tweet/internal/dto"
	"github.com/Khatchi/go-tweet/internal/model"
)

func (s *postService) CreatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, userID int64) (int64, int, error) {
	// store post
	now := time.Now()
	insertedID, err := s.postRepo.StorePost(ctx, &model.PostModel{
		UserID:    userID,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	// return
	return insertedID, http.StatusCreated, nil

}
