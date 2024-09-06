package svc

import (
	"github.com/Sion-L/devops/user/internal/config"
	"github.com/Sion-L/devops/user/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	AuthUsersModel    model.AuthUsersModel
	SettingsModel     model.SettingsModel
	AuthRoleMenuModel model.AuthRoleMenuModel
	AuthMenusModel    model.AuthMenusModel
	RedisClient       *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:            c,
		AuthUsersModel:    model.NewAuthUsersModel(conn, c.Cache),
		SettingsModel:     model.NewSettingsModel(conn, c.Cache),
		AuthRoleMenuModel: model.NewAuthRoleMenuModel(conn, c.Cache),
		AuthMenusModel:    model.NewAuthMenusModel(conn, c.Cache),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type // 设置redis类型
			r.Pass = c.Redis.Pass
		}),
	}
}
