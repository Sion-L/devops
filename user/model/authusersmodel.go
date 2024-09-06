package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AuthUsersModel = (*customAuthUsersModel)(nil)

type (
	// AuthUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthUsersModel.
	AuthUsersModel interface {
		authUsersModel
	}

	customAuthUsersModel struct {
		*defaultAuthUsersModel
	}
)

// NewAuthUsersModel returns a model for the database table.
func NewAuthUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AuthUsersModel {
	return &customAuthUsersModel{
		defaultAuthUsersModel: newAuthUsersModel(conn, c, opts...),
	}
}

func (m *defaultAuthUsersModel) FindOneByUser(ctx context.Context, user string) (*AuthUsers, error) {
	var resp AuthUsers
	query := "select %s from %s WHERE `username` = ?"
	sqlJoint := fmt.Sprintf(query, authUsersRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, sqlJoint, user)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
