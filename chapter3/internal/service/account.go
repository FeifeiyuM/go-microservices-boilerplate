package service

import (
	"context"
	"gmb/internal/model"
	"gmb/internal/dao"
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

func (srv *AccountSrv) saveAccount(ctx context.Context, acc *model.Account) error {
	return dao.GetAccountRepo().CreateAccount(ctx, acc)
}


func (srv *AccountSrv) CreateAccount(ctx context.Context, name, address string, gender int8) error {
	acc := srv.newAccount(name, address, gender)
	return srv.saveAccount(ctx, acc)
}

