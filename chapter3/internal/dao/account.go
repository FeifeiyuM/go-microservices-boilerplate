package dao

import (
	"context"
	"gmb/internal/model"
)

type AccountRepo interface {
	// 新建账户
	CreateAccount(ctx context.Context, acc *model.Account) error
	// 获取账户
	GetAccountById(ctx context.Context, id int64) (*model.Account, error)
}

var accountRepo AccountRepo

type aRepoImpl struct {
	*DaoBase
}

// 初始化 account dao 层
func InitAccountRepo() {
	accountRepo = &aRepoImpl{
		DaoBase: daoBase,
	}
}

// 获取 account repo
func GetAccountRepo() AccountRepo {
	return accountRepo
}

// 新建账户
func (a *aRepoImpl) CreateAccount(ctx context.Context, acc *model.Account) error {
	if acc == nil {
		return nil
	}
	return a.d.Save(ctx, acc)
}

// 获取账户
func (a *aRepoImpl) GetAccountById(ctx context.Context, id int64) (*model.Account, error) {
	_, err := a.d.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	acc := &model.Account{
		Id: id,
		Name: "test",
	}
	return acc, nil
}
