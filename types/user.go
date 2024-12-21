package types

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserRequest struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	GrantType    string `json:"grantType"`
	RefreshToken string `json:"refreshToken"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

var (
	uppercaseRegex   = regexp.MustCompile(`[A-Z]`)                                          // At least one uppercase letter
	lowercaseRegex   = regexp.MustCompile(`[a-z]`)                                          // At least one lowercase letter
	numberRegex      = regexp.MustCompile(`[0-9]`)                                          // At least one number
	specialCharRegex = regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':"\\|,.<>\/?~]`) // At least one special character
)

func (user UserRequest) ValidateUserSignUp() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name,
			validation.Length(1, 50)),
		validation.Field(&user.Email,
			validation.Required.Error("email cannot be empty"),
			is.Email,
		),
		validation.Field(&user.Password,
			validation.Required.Error("password cannot be empty"),
			validation.Length(8, 100).Error("password must be between 8 and 100 characters"),
			validation.Match(uppercaseRegex).Error("password must contain at least one uppercase letter"),
			validation.Match(lowercaseRegex).Error("password must contain at least one lowercase letter"),
			validation.Match(numberRegex).Error("password must contain at least one number"),
			validation.Match(specialCharRegex).Error("password must contain at least one special character"),
		),
		validation.Field(&user.Role,
			validation.Required,
			validation.In("admin", "user"),
		),
	)
}

func (user LoginRequest) ValidateLogin() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email,
			validation.Required.Error("email cannot be empty"),
			is.Email,
		),
		validation.Field(&user.Password,
			validation.Required.Error("password cannot be empty"),
			validation.Length(8, 100).Error("password must be between 8 and 100 characters"),
			validation.Match(uppercaseRegex).Error("password must contain at least one uppercase letter"),
			validation.Match(lowercaseRegex).Error("password must contain at least one lowercase letter"),
			validation.Match(numberRegex).Error("password must contain at least one number"),
			validation.Match(specialCharRegex).Error("password must contain at least one special character"),
		),
		validation.Field(&user.GrantType,
			validation.Required,
			validation.In("password", "refreshToken"),
		),
	)
}
