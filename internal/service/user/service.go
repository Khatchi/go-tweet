package user

import (
	"github.com/Khatchi/go-tweet/internal/config"
	"github.com/Khatchi/go-tweet/internal/repository/user"
)

type UserService interface {
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
