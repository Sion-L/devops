package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AuthMenusModel = (*customAuthMenusModel)(nil)

type (
	// AuthMenusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthMenusModel.
	AuthMenusModel interface {
		authMenusModel
	}

	customAuthMenusModel struct {
		*defaultAuthMenusModel
	}
)

// NewAuthMenusModel returns a model for the database table.
func NewAuthMenusModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AuthMenusModel {
	return &customAuthMenusModel{
		defaultAuthMenusModel: newAuthMenusModel(conn, c, opts...),
	}
}
