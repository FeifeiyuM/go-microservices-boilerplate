package dao

import (
	"context"
	"fmt"
	"gmb/pkg/log"
)

// 数据库 mock
type DataBaseMock struct {}

func NewNewDataBase() *DataBaseMock {
	return &DataBaseMock{}
}

func (d *DataBaseMock) Save(_ context.Context, data interface{}) error {
	fmt.Printf("Saved Data: %v", data)
	return nil
}

func (d *DataBaseMock) Get(_ context.Context, param interface{}) (interface{}, error) {
	fmt.Printf("Get Data: %v", param)
	return nil, nil
}

// dao 层基本对象
type DaoBase struct {
	d *DataBaseMock
	logger log.Factory
}

var daoBase *DaoBase

func InitDao(dataBase *DataBaseMock, logger log.Factory) {
	daoBase = &DaoBase{}
}

