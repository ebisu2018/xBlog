package user

import (
	"encoding/json"

	"github.com/ebisu2018/xBlog/common"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(req *CreateUserRequest) *User {
	req.PasswordCrypto()
	return &User{
		common.NewMetaData(),
		req,
	}
}

type User struct {
	*common.MetaData
	*CreateUserRequest
}

func (u *User) String() string {
	b, _ := json.Marshal(u)
	return string(b)
}

func (u *User) TableName() string {
	return "users"
}

func (u *User)CryptoCheck(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}