// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	settingsFieldNames          = builder.RawFieldNames(&Settings{})
	settingsRows                = strings.Join(settingsFieldNames, ",")
	settingsRowsExpectAutoSet   = strings.Join(stringx.Remove(settingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	settingsRowsWithPlaceHolder = strings.Join(stringx.Remove(settingsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheSettingsIdPrefix   = "cache:settings:id:"
	cacheSettingsNamePrefix = "cache:settings:name:"
)

type (
	settingsModel interface {
		Insert(ctx context.Context, data *Settings) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Settings, error)
		FindOneByName(ctx context.Context, name string) (*Settings, error)
		Update(ctx context.Context, data *Settings) error
		Delete(ctx context.Context, id int64) error
		FindLdapSettings(ctx context.Context, settings []string) ([]map[string]string, error)
	}

	defaultSettingsModel struct {
		sqlc.CachedConn
		table string
	}

	Settings struct {
		Id        int64  `db:"id"`
		Name      string `db:"name"`
		Value     string `db:"value"`
		Category  string `db:"category"`  // 设置类别
		Encrypted int64  `db:"encrypted"` // value是否加密
	}
)

func newSettingsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSettingsModel {
	return &defaultSettingsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`settings`",
	}
}

func (m *defaultSettingsModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	settingsIdKey := fmt.Sprintf("%s%v", cacheSettingsIdPrefix, id)
	settingsNameKey := fmt.Sprintf("%s%v", cacheSettingsNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, settingsIdKey, settingsNameKey)
	return err
}

func (m *defaultSettingsModel) FindOne(ctx context.Context, id int64) (*Settings, error) {
	settingsIdKey := fmt.Sprintf("%s%v", cacheSettingsIdPrefix, id)
	var resp Settings
	err := m.QueryRowCtx(ctx, &resp, settingsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", settingsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSettingsModel) FindOneByName(ctx context.Context, name string) (*Settings, error) {
	settingsNameKey := fmt.Sprintf("%s%v", cacheSettingsNamePrefix, name)
	var resp Settings
	err := m.QueryRowIndexCtx(ctx, &resp, settingsNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", settingsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSettingsModel) Insert(ctx context.Context, data *Settings) (sql.Result, error) {
	settingsIdKey := fmt.Sprintf("%s%v", cacheSettingsIdPrefix, data.Id)
	settingsNameKey := fmt.Sprintf("%s%v", cacheSettingsNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, settingsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Value, data.Category, data.Encrypted)
	}, settingsIdKey, settingsNameKey)
	return ret, err
}

func (m *defaultSettingsModel) Update(ctx context.Context, newData *Settings) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	settingsIdKey := fmt.Sprintf("%s%v", cacheSettingsIdPrefix, data.Id)
	settingsNameKey := fmt.Sprintf("%s%v", cacheSettingsNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, settingsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.Value, newData.Category, newData.Encrypted, newData.Id)
	}, settingsIdKey, settingsNameKey)
	return err
}

func (m *defaultSettingsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSettingsIdPrefix, primary)
}

func (m *defaultSettingsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", settingsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSettingsModel) tableName() string {
	return m.table
}
