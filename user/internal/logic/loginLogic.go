package logic

import (
	"context"
	"errors"
	"fmt"
	pkg "github.com/Sion-L/devops/pkg/user"
	"github.com/Sion-L/devops/user/internal/svc"
	"github.com/Sion-L/devops/user/model"
	"github.com/Sion-L/devops/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	users, err := l.svcCtx.AuthUsersModel.FindOneByUser(l.ctx, in.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			// 用户不存在
			return nil, fmt.Errorf("failed to find user: %w", err)
		}
		return &user.LoginResp{}, err
	}

	if users.Source == "local" {
		return l.LocalLogin(users, in.Password)
	}

	if users.Source == "ldap" {
		return l.LdapLogin(in, users)
	}

	return &user.LoginResp{}, nil
}

func (l *LoginLogic) LocalLogin(users *model.AuthUsers, password string) (*user.LoginResp, error) {
	if password == users.Password {
		return &user.LoginResp{
			UserId:   users.UserId,
			Username: users.Username,
			Token:    "",         // Token 应该由实际的认证逻辑生成
			ExpireAt: 2356782323, // 到期时间应根据需求生成
		}, nil
	}
	return &user.LoginResp{}, fmt.Errorf("failed to authenticate user: %s", users.Username)
}

func (l *LoginLogic) LdapLogin(in *user.LoginReq, users *model.AuthUsers) (*user.LoginResp, error) {
	ldapSettings, err := l.svcCtx.SettingsModel.FindLdapSettings(l.ctx)
	if err != nil {
		return &user.LoginResp{}, fmt.Errorf("failed to load LDAP settings: %w", err)
	}
	ldap := &pkg.LDAPServer{
		ServerUrl:  ldapSettings[pkg.LDAP_URL],
		BaseDN:     ldapSettings[pkg.LDAP_BINDDN],
		BindDN:     ldapSettings[pkg.LDAP_OU],
		BindPass:   ldapSettings[pkg.LDAP_PASSWORD],
		UserFilter: ldapSettings[pkg.LDAP_FILTER],
		UserAttr:   ldapSettings[pkg.LDAP_USERATTR],
	}

	if err := ldap.VerifyLDAPUser(in.Username, in.Password); err != nil {
		return &user.LoginResp{}, err
	}
	return &user.LoginResp{
		UserId:   users.UserId,
		Username: users.Username,
		Token:    "",
		ExpireAt: 0,
	}, nil
}
