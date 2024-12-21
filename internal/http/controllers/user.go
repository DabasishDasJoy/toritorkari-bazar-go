package controllers

import (
	"net/http"
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/types"

	"github.com/labstack/echo/v4"
)

var UserService domain.IUserService

func SetUserServiceInstance(userService domain.IUserService) {
	UserService = userService
}

func SignUp(e echo.Context) error {
	reqUser := types.UserRequest{}

	if err := e.Bind(&reqUser); err != nil {
		return e.JSON(http.StatusBadRequest, "Error binding")
	}

	if err := reqUser.ValidateUserSignUp(); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := UserService.SignUp(reqUser); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, "successfully signed up")
}

func Login(e echo.Context) error {
	loginRequest := types.LoginRequest{}

	if err := e.Bind(&loginRequest); err != nil {
		return e.JSON(http.StatusBadRequest, "error binding login request")
	}

	if err := loginRequest.ValidateLogin(); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var (
		loginResponse types.LoginResponse
		err           error
	)

	if loginRequest.GrantType == "password" {
		loginResponse, err = UserService.Login(loginRequest)
	} else if loginRequest.GrantType == "refreshToken" {
		loginResponse, err = UserService.RefreshToken(loginRequest)
	}

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return e.JSON(http.StatusFound, loginResponse)
}
