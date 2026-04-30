package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Khatchi/go-tweet/internal/dto"
	"github.com/Khatchi/go-tweet/internal/model"
	"github.com/Khatchi/go-tweet/pkg/jwt"
	"github.com/Khatchi/go-tweet/pkg/refreshtoken"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) {
	// checks if user exists
	userExists, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, "")
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExists == nil {
		return "", "", http.StatusNotFound, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(req.Password))
	if err != nil {
		return "", "", http.StatusNotFound, errors.New("invalid credentials")
	}

	// generate access token
	token, err := jwt.CreateToken(userExists.ID, userExists.UserName, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// get refresh token if exists
	now := time.Now()
	refreshTokenExists, err := s.userRepo.GetRefreshToken(ctx, userExists.ID, now)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExists != nil {
		return token, refreshTokenExists.RefreshToken, http.StatusOK, nil
	}

	// generate & refresh token if it doesn't exist
	refreshToken, err := refreshtoken.GetRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	fmt.Printf("DEBUG: userExists.ID = %d\n", userExists.ID)
	fmt.Printf("DEBUG: userExists = %+v\n", userExists)

	err = s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:       userExists.ID,
		RefreshToken: refreshToken,
		CreatedAt:    now,
		UpdatedAt:    now,
		ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
	})

	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// return
	return token, refreshToken, http.StatusOK, nil
}
