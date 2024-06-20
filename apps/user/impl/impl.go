package impl

import (
	"context"

	"github.com/ebisu2018/xBlog/apps/user"
	"github.com/ebisu2018/xBlog/common"
	"github.com/ebisu2018/xBlog/config"
	"github.com/ebisu2018/xBlog/exception"
	"gorm.io/gorm"
)


var (
	_ user.UserService = &UserServiceImpl{}
)


func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		cfg: config.ReadConfig(),
	}
}


type UserServiceImpl struct {
	cfg *config.Config
}


func (i *UserServiceImpl)CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error)  {
	if err := req.ValidateAccount(); err != nil {
		return nil, err
	}

	ins := user.NewUser(req)
	err := i.cfg.MySql.GetConn().WithContext(ctx).Create(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}


func (i *UserServiceImpl)DeleteUser(ctx context.Context, req *user.DeleteUserRequest) error {

	ins, err := i.QueryUser(ctx, user.NewQueryRequestId(common.ParseInt(req.Id)))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exception.NewRecordNotFound("user not found")
		}
	}
	return i.cfg.MySql.GetConn().WithContext(ctx).Where("id = ?", req.Id).Delete(ins).Error
}


func (i *UserServiceImpl)QueryUser(ctx context.Context, req *user.QueryRequest) (*user.User, error)  {
	query := i.cfg.MySql.GetConn().WithContext(ctx)
	switch req.QueryBy {
	case user.QueryById:
		query = query.Where("id = ?", req.QueryValue)
	case user.QueryByName:
		query = query.Where("username = ?", req.QueryValue)
	}

	ins := user.NewUser(user.NewCreateUserRequest())
	err := query.First(ins).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewRecordNotFound("user not found")
		}
	}
	return ins, nil
}
