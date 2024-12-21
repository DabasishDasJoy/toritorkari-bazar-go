package domain

import (
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type IUserRepo interface {
	SignUp(user models.User) error
	GetUserByEmail(email string) (models.User, error)
}

type IUserService interface {
	SignUp(user types.UserRequest) error
	GetUserByEmail(email string) (types.UserRequest, error)
	Login(loginRequest types.LoginRequest) (types.LoginResponse, error)
	RefreshToken(loginRequest types.LoginRequest) (types.LoginResponse, error)
}
