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
)
