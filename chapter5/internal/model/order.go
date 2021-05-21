package model

import (
	"strconv"
	"time"
)

type Order struct {
	Id int64 `json:"id" sql:"id"`
	// 账号id
	AccId int64 `json:"acc_id" sql:"acc_id"`
	// 订单号
	OrderNumber string `json:"order_number" sql:"order_number"`
	// 外部支付单号
	PayOrderId string `json:"pay_order_id" sql:"pay_order_id"`
	// 订单接口(精确至分)
	Amount int64 `json:"amount" sql:"amount"`
	// 创建时间
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at" sql:"updated_at"`
	// 删除时间
	DeletedAt int64 `json:"deleted_at" sql:"deleted_at"`
}

func NewOrder(accId, Amt int64, payOrderId string) *Order {
	now := time.Now()
	o := &Order{
		AccId:       accId,
		OrderNumber: "",
		Amount:      Amt,
		PayOrderId:  payOrderId,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	o.GenOrderNum()
	return o
}

func (o *Order) GenOrderNum() {
	ts := o.CreatedAt.Nanosecond()
	o.OrderNumber = strconv.FormatInt(int64(ts), 10)
}
