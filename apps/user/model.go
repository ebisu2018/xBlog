package user

import (
	"encoding/json"

	"github.com/ebisu2018/xBlog/common"
)

func NewUser(req *CreateUserRequest) *User {
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

func (u *User)TableName() string {
	return "users"
}