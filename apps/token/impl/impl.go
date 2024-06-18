package impl

import (
	"context"

	"github.com/ebisu2018/xBlog/apps/token"
	"github.com/ebisu2018/xBlog/config"
)

func NewTokenServiceImpl() *TokenServiceImpl {
	return &TokenServiceImpl{
		cfg: config.ReadConfig(),
	}
}

type TokenServiceImpl struct {
	cfg *config.Config
}

func (i *TokenServiceImpl)Login(ctx context.Context, req *token.LoginRequest) (*token.Token, error)  {
	return nil, nil
}

func (i *TokenServiceImpl)Validate(ctx context.Context, req *token.ValidateRequest) error  {
	return nil
}

func (i *TokenServiceImpl)Logout(ctx context.Context, req *token.LogoutRequest) error  {
	return nil
}