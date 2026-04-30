package model

import "time"

type (
	UserModel struct {
		ID        int64
		Email     string
		UserName  string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	RefreshTokenModel struct {
		ID           int64
		UserID       int64
		RefreshToken string
		ExpiredAt    time.Time
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)
