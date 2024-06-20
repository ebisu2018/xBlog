package token

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/rs/xid"
)

func NewToken() *Token {
	return &Token{
		AccessToken: xid.New().String(),
		AccessTokenExpiredAt: 7200,
		RefreshToken: xid.New().String(),
		RefreshTokenExpiredAt: 3600 * 24 * 7,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}

type Token struct {
	UserId                int    `json:"id" gorm:"column:user_id"`
	UserName              string `json:"username"`
	CreatedAt             int64  `json:"created_at"`
	UpdatedAt             int64  `json:"updated_at"`
	AccessToken           string `json:"access_token"`
	AccessTokenExpiredAt  int64  `json:"access_token_expired_at"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiredAt int64  `json:"refresh_token_expired_at"`
}

func (t *Token) String() string {
	data, _ := json.Marshal(t)
	return string(data)
}

func (t *Token)TableName() string {
	return "tokens"
}


func (t *Token)expiredTime() time.Time {
	expiredTime := time.Unix(t.CreatedAt, 0).Add(time.Duration(t.AccessTokenExpiredAt) * time.Second)
	fmt.Printf("expired Time is %v\n", expiredTime)
	return expiredTime
}


func (t *Token)IsExpired() error {
	timeout := time.Since(t.expiredTime()).Seconds()
	fmt.Printf("timeout for %v seconds\n", timeout)
	if timeout > 0 {
		return fmt.Errorf(" timeout %v seconds", timeout)
	}
	return nil
}
