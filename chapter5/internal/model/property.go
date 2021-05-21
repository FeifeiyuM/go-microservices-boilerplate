package model

import (
	"time"
)

type Property struct {
	Id int64 `json:"id" sql:"id"`
	// 账号id
	AccId int64 `json:"acc_id" sql:"acc_id"`
	// 余额(精确至分）
	Balance int64 `json:"balance" sql:"balance"`
	// 创建时间
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at" sql:"updated_at"`
	// 删除时间
	DeletedAt int64 `json:"deleted_at" sql:"deleted_at"`
}

func NewProperty(accId int64, balance int64) *Property {
	return &Property{
		AccId:     accId,
		Balance:   balance,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: 0,
	}
}

func (p *Property) AddBalance(balance int64) {
	p.Balance += balance
	p.UpdatedAt = time.Now()
}
