package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SettingsModel = (*customSettingsModel)(nil)

type (
	// SettingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSettingsModel.
	SettingsModel interface {
		settingsModel
	}

	customSettingsModel struct {
		*defaultSettingsModel
	}
)

// NewSettingsModel returns a model for the database table.
func NewSettingsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SettingsModel {
	return &customSettingsModel{
		defaultSettingsModel: newSettingsModel(conn, c, opts...),
	}
}

func (m *defaultSettingsModel) FindLdapSettings(ctx context.Context, settings []string) ([]map[string]string, error) {

	result := make([]map[string]string, 10)
	for _, setting := range settings {
		settingsNameKey := fmt.Sprintf("%s%v", cacheSettingsNamePrefix, setting)
		var resp Settings
		err := m.QueryRowIndexCtx(ctx, &resp, settingsNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
			query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", settingsRows, m.table)
			if err := conn.QueryRowCtx(ctx, &resp, query, setting); err != nil {
				return nil, err
			}
			return resp.Id, nil
		}, m.queryPrimary)
		switch err {
		case nil:
			result = append(result, map[string]string{
				setting: resp.Value,
			})
		case sqlc.ErrNotFound:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	logx.Info("result: %v", result)
	return result, nil
}
