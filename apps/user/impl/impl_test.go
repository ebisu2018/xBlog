package impl_test

import (
	"context"
	"testing"

	"github.com/ebisu2018/xBlog/apps/user"
	"github.com/ebisu2018/xBlog/common"
	"github.com/ebisu2018/xBlog/config"
	"github.com/ebisu2018/xBlog/ioc"
	_ "github.com/ebisu2018/xBlog/apps"
)

var (
	// userImpl = impl.NewUserServiceImpl()
	userImpl user.UserService
	ctx      = context.Background()
)

func init() {
	config.LoadFromTomlFile()
	ioc.Container().InitObj()
	userImpl = ioc.Container().Get(user.AppName).(user.UserService)
}

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.UserName = "testuser"
	req.Password = "12345"
	ins, err := userImpl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDeleteUser(t *testing.T) {
	req := user.NewDeleteUserRequest(1)
	err := userImpl.DeleteUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestQueryUserById(t *testing.T) {

	req := user.NewQueryRequestId(common.ParseInt(6))
	ins, err := userImpl.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryUserByName(t *testing.T) {
	req := user.NewQueryRequestName("testuser")
	ins, err := userImpl.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
