package model

import (
	"time"

	"github.com/FeifeiyuM/sqly"
)

const (
	UnKnow int8 = 0
	Female int8 = 1
	Male   int8 = 2
)

// Account
type Account struct {
	Id int64 `json:"id" sql:"id"`
	// 昵称
	Nickname string `json:"nickname" sql:"nickname"`
	// 头像地址
	Avatar sqly.NullString `json:"avatar" sql:"avatar"`
	// 性别
	Gender int8 `json:"gender" sql:"gender"`
	// 生日
	Birthday sqly.NullTime `json:"birthday" sql:"birthday"`
	// 手机号
	Mobile string `json:"mobile" sql:"mobile"`
	// 密码
	Password string `json:"password" sql:"password"`
	// 创建时间
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at" sql:"updated_at"`
	// 删除时间
	DeletedAt int64 `json:"deleted_at" sql:"deleted_at"`
}

// 初始化 Account
func NewAccount(nickname, avatar, mobile string, birthday time.Time) *Account {
	a := sqly.NullString{}
	if avatar != "" {
		a = sqly.NullString{String: avatar, Valid: true}
	}
	b := sqly.NullTime{}
	if !birthday.IsZero() {
		b = sqly.NullTime{Time: birthday, Valid: true}
	}
	return &Account{
		Nickname:  nickname,
		Avatar:    a,
		Birthday:  b,
		Mobile:    mobile,
		CreatedAt: time.Now(),
	}
}

// 设置性别
func (a *Account) SetGender(g int8) {
	switch g {
	case Female, Male:
		a.Gender = g
	default:
		a.Gender = UnKnow
	}
}
