package impl

import (
	"context"
	"fmt"

	"github.com/ebisu2018/xBlog/apps/token"
	"github.com/ebisu2018/xBlog/apps/user"
	"github.com/ebisu2018/xBlog/config"
	"github.com/ebisu2018/xBlog/exception"
	"github.com/ebisu2018/xBlog/ioc"
	"gorm.io/gorm"
)

func init() {
	fmt.Println("loading token package")
	ioc.Container().Registry(&TokenServiceImpl{})
}

// var _ token.TokenService = NewTokenServiceImpl(impl.NewUserServiceImpl())

func (i *TokenServiceImpl) Init() {
	i.cfg = config.ReadConfig()
	i.userSvc = ioc.Container().Get(user.AppName).(user.UserService)
}

func (i *TokenServiceImpl) Name() string {
	return token.AppName
}

func NewTokenServiceImpl(svc user.UserService) *TokenServiceImpl {
	return &TokenServiceImpl{
		cfg:     config.ReadConfig(),
		userSvc: svc,
	}
}

type TokenServiceImpl struct {
	cfg     *config.Config
	userSvc user.UserService
}

func (i *TokenServiceImpl) Login(ctx context.Context, req *token.LoginRequest) (*token.Token, error) {

	// 查询是否有该用户
	u, err := i.userSvc.QueryUser(ctx, user.NewQueryRequestName(req.Username))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewRecordNotFound("requested user %v not found", req.Username)
		}
	}

	// 有则比对密钥是否正确
	if err := u.CryptoCheck(req.Password); err != nil {
		return nil, exception.NewAuthFailed("username or password is invalid")
	}

	// 创建token并写入库
	tk := token.NewToken()
	tk.UserId = u.Id
	tk.UserName = u.UserName

	err = i.cfg.MySql.GetConn().WithContext(ctx).Create(tk).Error
	if err != nil {
		return nil, err
	}
	return tk, nil
}

func (i *TokenServiceImpl) Validate(ctx context.Context, req *token.ValidateRequest) error {

	tk := token.NewToken()
	err := i.cfg.MySql.GetConn().WithContext(ctx).Where("access_token = ?", req.AccessToken).First(tk).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exception.NewRecordNotFound("token %v not found", req.AccessToken)
		}
	}

	if err := tk.IsExpired(); err != nil {
		return exception.NewRequestTimeout("access token %v is expired", req.AccessToken)
	}
	return nil
}

func (i *TokenServiceImpl) Logout(ctx context.Context, req *token.LogoutRequest) error {
	return nil
}
