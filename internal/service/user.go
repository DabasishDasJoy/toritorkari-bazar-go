package service

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/methods"
	"toritorkari-bazar/types"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo domain.IUserRepo
}

func UserServiceInstance(userRepo domain.IUserRepo) domain.IUserService {
	return &UserService{
		repo: userRepo,
	}
}

func (service UserService) SignUp(user types.UserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	userRequest := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
		Role:     user.Role,
	}

	if err := service.repo.SignUp(userRequest); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("email already exists")
		}
		return err
	}

	return nil
}

func (service UserService) GetUserByEmail(email string) (types.UserRequest, error) {
	user, err := service.repo.GetUserByEmail(email)
	if err != nil {
		return types.UserRequest{}, err
	}

	return types.UserRequest{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		Password: user.Password,
	}, nil
}

func (service UserService) Login(loginRequest types.LoginRequest) (types.LoginResponse, error) {
	user, err := service.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return types.LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return types.LoginResponse{}, errors.New("invalid password")
	}

	accessToken, err := methods.GenerateToken(user.Email, 5*time.Minute)
	if err != nil {
		return types.LoginResponse{}, err
	}

	refreshToken, err := methods.GenerateToken(user.Email, 24*time.Hour)
	if err != nil {
		return types.LoginResponse{}, err
	}

	return types.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (service UserService) RefreshToken(loginRequest types.LoginRequest) (types.LoginResponse, error) {
	var (
		claims models.Claims
		err    error
	)

	if claims, err = methods.ValidateRefreshToken(loginRequest.RefreshToken); err != nil {
		return types.LoginResponse{}, err
	}

	accessToken, err := methods.GenerateToken(claims.Subject, 5*time.Minute)

	if err != nil {
		return types.LoginResponse{}, err
	}

	return types.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: loginRequest.RefreshToken,
	}, nil
}
