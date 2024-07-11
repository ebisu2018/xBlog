package impl_test

import (
	"context"
	"testing"

	"github.com/ebisu2018/xBlog/apps/token"
	// tokenImpl "github.com/ebisu2018/xBlog/apps/token/impl"
	// userImpl "github.com/ebisu2018/xBlog/apps/user/impl"
	"github.com/ebisu2018/xBlog/config"
	"github.com/ebisu2018/xBlog/ioc"
	_ "github.com/ebisu2018/xBlog/apps"
)

var (
	ctx = context.Background()
	// tkSvc = tokenImpl.NewTokenServiceImpl(userImpl.NewUserServiceImpl())
	tkSvc = ioc.Container().Get(token.AppName).(token.TokenService)
)


func init()  {
	config.LoadFromTomlFile()
	ioc.Container().InitObj()
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
	err := tkSvc.Validate(ctx, token.NewValidateRequest("cq7k1scb3284fu1pqctg"))
	if err != nil {
		t.Fatal(err)
	}
}