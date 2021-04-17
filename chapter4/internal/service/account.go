package service

import (
	"context"
	"errors"
	"gmb/internal/model"
	"gmb/internal/dao"
	"gmb/pkg/gmberror"
)

type AccountSrv struct {

}

var accSrv *AccountSrv
// 初始化
func InitAccountSrv() {
	accSrv = &AccountSrv{}
}
// 获取服务
func GetAccountSrv() *AccountSrv {
	return accSrv
}

// 新建 account
func (srv *AccountSrv) newAccount(name, address string, gender int8) *model.Account {
	return &model.Account{
		Name: name,
		Address: address,
		Gender: gender,
	}
}

func (srv *AccountSrv) saveAccount(ctx context.Context, acc *model.Account) gmberror.GMBError {
	err := dao.GetAccountRepo().CreateAccount(ctx, acc)
	if err != nil {
		return gmberror.DBError(err)
	}
	return nil
}


func (srv *AccountSrv) CreateAccount(ctx context.Context, name, address string, gender int8) gmberror.GMBError {
	if gender != model.Female && gender != model.Male {
		return gmberror.InvalidRequest(errors.New("性别错误"))
	}
	acc := srv.newAccount(name, address, gender)
	err := srv.saveAccount(ctx, acc)
	if err != nil {
		return gmberror.DBError(err)
	}
	return nil
}

