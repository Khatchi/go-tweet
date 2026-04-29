package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Khatchi/go-tweet/internal/dto"
	"github.com/Khatchi/go-tweet/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error) {
	//checks if user already exists
	userExist, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, req.Username)
	if userExist != nil {
		return 0, http.StatusBadRequest, errors.New("user already exist")
	}

	//hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	//create user
	now := time.Now()
	userModel := &model.UserModel{
		Email:     req.Email,
		UserName:  req.Username,
		Password:  string(passwordHash),
		CreatedAt: now,
		UpdatedAt: now,
	}

	userID, err := s.userRepo.CreateUser(ctx, userModel)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return userID, http.StatusCreated, nil

}
