package service

import (
	"context"
	"gmb/internal/conf"
	"gmb/internal/dao"
	"gmb/internal/model"
	"gmb/pkg/gmberror"
	"gmb/pkg/log"
	"go.uber.org/zap"
	"time"
)

type AccountSrv struct {
	cfg    *conf.Config
	logger log.Factory
}

var accSrv *AccountSrv

// 初始化
func InitAccountSrv(cfg *conf.Config, logger log.Factory) {
	accSrv = &AccountSrv{
		cfg:    cfg,
		logger: logger,
	}
}

// 获取服务
func GetAccountSrv() *AccountSrv {
	return accSrv
}

func (srv *AccountSrv) saveAccount(ctx context.Context, acc *model.Account) gmberror.GMBError {
	id, err := dao.GetAccountRepo().CreateAccount(ctx, acc)
	if err != nil {
		return gmberror.DBError(err)
	}
	acc.Id = id
	return nil
}

// 新建 account
func (srv *AccountSrv) CreateAccount(ctx context.Context, name, avatar, mobile string, gender int8) gmberror.GMBError {
	srv.logger.For(ctx).Info("create account", zap.String("name", name))
	acc := model.NewAccount(name, avatar, mobile, time.Time{})
	acc.SetGender(gender)
	err := srv.saveAccount(ctx, acc)
	if err != nil {
		return gmberror.DBError(err)
	}
	return nil
}
