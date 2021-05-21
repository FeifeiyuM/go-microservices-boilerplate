package dao

import (
	"github.com/FeifeiyuM/sqly"
	"github.com/go-redis/redis/v7"
	"gmb/pkg/log"
)

// dao 层基本对象
type DaoBase struct {
	db     *sqly.SqlY
	cache  *redis.Client
	logger log.Factory
}

var daoBase *DaoBase

func InitDao(db *sqly.SqlY, cache *redis.Client, logger log.Factory) {
	daoBase = &DaoBase{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}
