package impl_test

import (
	"context"
	"testing"

	"github.com/ebisu2018/xBlog/apps/token"
	tokenImpl "github.com/ebisu2018/xBlog/apps/token/impl"
	userImpl "github.com/ebisu2018/xBlog/apps/user/impl"
	"github.com/ebisu2018/xBlog/config"
)

var (
	ctx = context.Background()
	tkSvc = tokenImpl.NewTokenServiceImpl(userImpl.NewUserServiceImpl())
)


func init()  {
	config.LoadFromTomlFile()
}

func TestLogin(t *testing.T)  {
	req := token.NewLoginRequest()
	req.Username = "testuser"
	req.Password = "12345"
	tk, err := tkSvc.Login(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestValidate(t *testing.T)  {
	err := tkSvc.Validate(ctx, token.NewValidateRequest("cppr5b4b3280qo65u6ng"))
	if err != nil {
		t.Fatal(err)
	}
}