package impl_test

import (
	"context"
	"testing"

	"github.com/ebisu2018/xBlog/apps/user"
	"github.com/ebisu2018/xBlog/apps/user/impl"
	"github.com/ebisu2018/xBlog/common"
	"github.com/ebisu2018/xBlog/config"
)

var (
	userImpl = impl.NewUserServiceImpl()
	ctx      = context.Background()
)

func init() {
	config.LoadFromTomlFile()
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

	req := user.NewQueryRequestId(common.ParseInt(5))
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
