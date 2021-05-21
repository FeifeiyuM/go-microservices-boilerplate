package dao

import (
	"context"
	"gmb/internal/model"
)

type AccountRepo interface {
	// 新建账户
	CreateAccount(ctx context.Context, acc *model.Account) (int64, error)
	// 获取账户
	GetAccountById(ctx context.Context, id int64) (*model.Account, error)
}

var accountRepo AccountRepo

// 获取 account repo
func GetAccountRepo() AccountRepo {
	return accountRepo
}

type aRepoImpl struct {
	*DaoBase
	name string
}

// 初始化 account dao 层
func InitAccountRepo() {
	accountRepo = &aRepoImpl{
		DaoBase: daoBase,
		name:    "dao_account",
	}
}

// 新建账户
func (a *aRepoImpl) CreateAccount(ctx context.Context, acc *model.Account) (int64, error) {
	if acc == nil {
		return 0, nil
	}
	param := []interface{}{
		acc.Nickname,
		acc.Avatar,
		acc.Gender,
		acc.Mobile,
		acc.Password,
		acc.CreatedAt,
	}
	query := "INSERT INTO `account` (`nickname`, `avatar`, `birthday`, `gender`, `mobile`, `password`, `created_at`) VALUES " +
		"(?,?,?,?,?,?,?)"
	aff, err := a.db.InsertCtx(ctx, query, param...)
	if err != nil {
		return 0, err
	}
	lastId, _ := aff.GetLastId()
	return lastId, nil
}

// 获取账户
func (a *aRepoImpl) GetAccountById(ctx context.Context, id int64) (*model.Account, error) {
	acc := &model.Account{}
	query := "SELECT * FROM `account` WHERE `id`=?"
	err := a.db.Get(acc, query, id)
	return acc, err
}
