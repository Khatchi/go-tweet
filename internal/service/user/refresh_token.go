package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Khatchi/go-tweet/internal/dto"
	"github.com/Khatchi/go-tweet/internal/model"
	"github.com/Khatchi/go-tweet/pkg/jwt"
	"github.com/Khatchi/go-tweet/pkg/refreshtoken"
)

func (s *userService) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int64) (string, string, int, error) {
	// check if user exists
	userExists, err := s.userRepo.GetUSerByID(ctx, userID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExists == nil {
		return "", "", http.StatusNotFound, errors.New("user not found")
	}

	// get refresh token by user id
	refreshtokenExists, err := s.userRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		return "", "", http.StatusUnauthorized, errors.New("refresh token had expired")
	}

	// check refresh token matches with request body
	if req.RefreshToken != refreshtokenExists.RefreshToken {
		return "", "", http.StatusUnauthorized, errors.New("refresh token not found")
	}

	// generate new token
	token, err := jwt.CreateToken(userID, userExists.UserName, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// delete old refresh token & generate new refresh token
	err = s.userRepo.DeleteRefreshTokenByUserID(ctx, userID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	refreshtoken, err := refreshtoken.GetRefreshToken()
	if err != nil {
		return "", "", http.StatusNotFound, err
	}

	now := time.Now()
	s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:       userID,
		RefreshToken: refreshtoken,
		CreatedAt:    now,
		UpdatedAt:    now,
		ExpiredAt:    now.Add(7 * 24 * time.Hour),
	})

	return token, refreshtoken, http.StatusOK, nil
}
