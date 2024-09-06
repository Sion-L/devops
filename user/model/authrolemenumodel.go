package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AuthRoleMenuModel = (*customAuthRoleMenuModel)(nil)

type (
	// AuthRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthRoleMenuModel.
	AuthRoleMenuModel interface {
		authRoleMenuModel
	}

	customAuthRoleMenuModel struct {
		*defaultAuthRoleMenuModel
	}
)

// NewAuthRoleMenuModel returns a model for the database table.
func NewAuthRoleMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AuthRoleMenuModel {
	return &customAuthRoleMenuModel{
		defaultAuthRoleMenuModel: newAuthRoleMenuModel(conn, c, opts...),
	}
}
