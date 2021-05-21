package service

import (
	"context"
	"fmt"
	"gmb/internal/conf"
	"gmb/internal/dao"
	"gmb/internal/model"
	"gmb/pkg/gmberror"
	"gmb/pkg/log"
)

type OrderSrv struct {
	cfg    *conf.Config
	logger log.Factory
}

var orderSrv *OrderSrv

// 初始化
func InitOrderSrv(cfg *conf.Config, logger log.Factory) {
	orderSrv = &OrderSrv{
		cfg:    cfg,
		logger: logger,
	}
}

// 获取服务
func GetOrderSrv() *OrderSrv {
	return orderSrv
}

// 账户检查
func (srv *OrderSrv) checkAccount(ctx context.Context, accId int64) (*model.Property, gmberror.GMBError) {
	// 账户检查
	acc, err := dao.GetAccountRepo().GetAccountById(ctx, accId)
	if err != nil {
		return nil, gmberror.DBError(err)
	}
	if acc == nil {
		return nil, gmberror.InvalidAccount(fmt.Errorf("accout: %d not available", accId))
	}
	// 账户资产信息检查
	pro, err := dao.GetPropertyRepo().GetPropertyByAccId(ctx, acc.Id)
	if err != nil {
		return nil, gmberror.DBError(err)
	}
	if pro == nil {
		return nil, gmberror.InvalidAccount(fmt.Errorf("account: %d property not available", accId))
	}
	return pro, nil
}

// 创建订单
func (srv *OrderSrv) createOrder(ctx context.Context, accId, amount int64, payOrderId string) (*model.Order, gmberror.GMBError) {
	// 检查订单是否已经存在，（假设我们可以根据支付单判断订单是否重复
	order, err := dao.GetOrderRepo().GetOrderByPayOrder(ctx, accId, payOrderId)
	if err != nil {
		return nil, gmberror.DBError(err)
	}
	if order != nil {
		return nil, gmberror.InvalidOrder(fmt.Errorf("order has created"))
	}
	// 生成订单
	order = model.NewOrder(accId, amount, payOrderId)
	order.GenOrderNum()
	order.Id, err = dao.GetOrderRepo().CreateOrder(ctx, order)
	if err != nil {
		return nil, gmberror.DBError(err)
	}
	return order, nil
}

func (srv *OrderSrv) updateProperty(ctx context.Context, pro *model.Property, amt int64) gmberror.GMBError {
	pro.AddBalance(amt)
	err := dao.GetPropertyRepo().UpdateProperty(ctx, pro)
	if err != nil {
		return gmberror.DBError(err)
	}
	return nil
}

// 账户充值
func (srv *OrderSrv) AccountRecharge(ctx context.Context, accId, amount int64, payOrderId string) gmberror.GMBError {

	// TODO 开启事务
	// 1、 检查账户
	pro, GErr := srv.checkAccount(ctx, accId)
	if GErr != nil {
		return GErr
	}
	// 2、检查并创建订单
	_, GErr = srv.createOrder(ctx, pro.AccId, amount, payOrderId)
	if GErr != nil {
		return GErr
	}
	// 更新资产
	GErr = srv.updateProperty(ctx, pro, amount)
	if GErr != nil {
		return GErr
	}
	// TODO 结束事务
	return nil
}
