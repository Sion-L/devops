package model

import (
	"context"
	"errors"
	"fmt"

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

func (m *defaultSettingsModel) FindLdapSettings(ctx context.Context) (map[string]string, error) {
	var settings []Settings
	result := make(map[string]string)
	query := "SELECT name,value FROM %s WHERE name IN ('AUTH_LDAP_SERVER_URI', 'AUTH_LDAP_BIND_PASSWORD', 'AUTH_LDAP_BIND_DN', 'AUTH_LDAP_SEARCH_FILTER', 'AUTH_LDAP_SEARCH_OU', 'AUTH_LDAP_SEARCH_FILTER')"
	sqlJoint := fmt.Sprintf(query, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &settings, sqlJoint)
	switch {
	case err == nil:
		for _, setting := range settings {
			result[setting.Name] = setting.Value
		}
		return result, nil
	case errors.Is(err, sqlc.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}

}
