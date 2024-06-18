package token

import "context"

type TokenService interface {
	Login(context.Context, *LoginRequest) (*Token, error)
	Logout(context.Context, *LogoutRequest) error
	Validate(context.Context, *ValidateRequest) error
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func NewValidateRequest(at string) *ValidateRequest {
	return &ValidateRequest{
		AccessToken: at,
	}	
}

type ValidateRequest struct {
	AccessToken string `json:"access_token"`
}


func NewLogoutRequest(at, rt string) *LogoutRequest {
	return &LogoutRequest{
		AccessToken: at,
		RefreshToken: rt,
	}
}

type LogoutRequest struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}