package user

import (
	"context"

	"github.com/Khatchi/go-tweet/internal/config"
	"github.com/Khatchi/go-tweet/internal/dto"
	"github.com/Khatchi/go-tweet/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error)
	Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) //main token, refrsh token, statuscode, error
	RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int64) (string, string, int, error)
}

type userService struct {
	cfg      *config.Config
	userRepo user.UserRepository
}

func NewService(cfg *config.Config, userRepo user.UserRepository) UserService {
	return &userService{
		cfg:      cfg,
		userRepo: userRepo,
	}
}
