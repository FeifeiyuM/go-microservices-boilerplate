package dao

import (
	"context"
	"gmb/internal/model"
)

type OrderRepo interface {
	// 新建订单
	CreateOrder(ctx context.Context, order *model.Order) (int64, error)
	// 根据支付单号获取订单
	GetOrderByPayOrder(ctx context.Context, accId int64, payOrderId string) (*model.Order, error)
}

var orderRepo OrderRepo

// 获取 order repo
func GetOrderRepo() OrderRepo {
	return orderRepo
}

type orderImpl struct {
	*DaoBase
	name string
}

// 初始化 account dao 层
func InitOrderRepo() {
	orderRepo = &orderImpl{
		DaoBase: daoBase,
		name:    "dao_order",
	}
}

// 新建订单
func (dao *orderImpl) CreateOrder(ctx context.Context, order *model.Order) (int64, error) {
	// TODO 实现创建订单
	return 0, nil
}

// 根据支付单号获取订单
func (dao *orderImpl) GetOrderByPayOrder(ctx context.Context, accId int64, payOrderId string) (*model.Order, error) {
	// TODO 实现获取订单
	return &model.Order{}, nil
}
