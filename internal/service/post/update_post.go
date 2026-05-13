package post

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Khatchi/go-tweet/internal/dto"
	"github.com/Khatchi/go-tweet/internal/model"
)

func (s *postService) UpdatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, postID int64, userID int64) (int, error) {
	// check if post exists
	postExists, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if postExists == nil {
		return http.StatusNotFound, errors.New("posts not found")
	}

	if postExists.UserID != userID {
		return http.StatusNotFound, errors.New("post not found")
	}

	// update post if exists
	err = s.postRepo.UpdatePost(ctx, &model.PostModel{
		Title:     req.Title,
		Content:   req.Content,
		UpdatedAt: time.Now(),
	}, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// return
	return http.StatusOK, nil
}
