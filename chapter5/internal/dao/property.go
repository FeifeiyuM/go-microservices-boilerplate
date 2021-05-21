package dao

import (
	"context"
	"gmb/internal/model"
)

type PropertyRepo interface {
	// 新建资产账户
	CreateProperty(ctx context.Context, pro *model.Property) (int64, error)
	// 获取账户
	GetPropertyByAccId(ctx context.Context, accId int64) (*model.Property, error)
	// 更新账户
	UpdateProperty(ctx context.Context, pro *model.Property) error
}

var propertyRepo PropertyRepo

// 获取 Property repo
func GetPropertyRepo() PropertyRepo {
	return propertyRepo
}

type pRepoImpl struct {
	*DaoBase
	name string
}

// 初始化 Property dao 层
func InitPropertyRepo() {
	propertyRepo = &pRepoImpl{
		daoBase,
		"dao_property",
	}
}

// 新建资产账户
func (dao *pRepoImpl) CreateProperty(ctx context.Context, pro *model.Property) (int64, error) {
	// TODO  实现创建账户
	return 0, nil
}

// 获取账户
func (dao *pRepoImpl) GetPropertyByAccId(ctx context.Context, accId int64) (*model.Property, error) {
	// TODO, 实现获取资产账户
	return &model.Property{}, nil
}

// 更新账户
func (dao *pRepoImpl) UpdateProperty(ctx context.Context, pro *model.Property) error {
	// TODO, 实现更新账户
	return nil
}
