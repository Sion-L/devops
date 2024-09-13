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
	u, err := l.svcCtx.AuthUsersModel.FindOneByUser(l.ctx, in.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			// 用户不存在
			return &user.LoginResp{}, fmt.Errorf("user not found: %s", in.Username)
		}
		return &user.LoginResp{}, err
	}

	l.Logger.Infof("用户信息: %v", u)
	if u.Source == "local" {
		return l.LocalLogin(u, in.Password)
	}

	if u.Source == "ldap" {
		return l.LdapLogin(in, u)
	}

	if err != nil {
		return nil, err
	}
	return &user.LoginResp{}, nil
}

func (l *LoginLogic) LocalLogin(u *model.AuthUsers, password string) (*user.LoginResp, error) {
	if password == u.Password {
		return &user.LoginResp{
			UserId:   u.UserId,
			Username: u.Username,
			RoleType: u.RoleType,
		}, nil
	}
	return &user.LoginResp{}, fmt.Errorf("failed to authenticate user: %s", u.Username)
}

func (l *LoginLogic) LdapLogin(in *user.LoginReq, u *model.AuthUsers) (*user.LoginResp, error) {
	ldap, err := ParseLdapConn(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	if err = ldap.VerifyLDAPUser(in.Username, in.Password); err != nil {
		return &user.LoginResp{}, err
	}
	return &user.LoginResp{
		UserId:   u.UserId,
		Username: u.Username,
		RoleType: u.RoleType,
	}, nil
}

func ParseLdapConn(ctx context.Context, svcCtx *svc.ServiceContext) (pkg.LDAPServer, error) {
	var ldap pkg.LDAPServer
	settings := []string{
		pkg.LDAP_URL,
		pkg.LDAP_OU,
		pkg.LDAP_BINDDN,
		pkg.LDAP_PASSWORD,
		pkg.LDAP_FILTER,
		pkg.LDAP_USERATTR}

	ldapSettings, err := svcCtx.SettingsModel.FindLdapSettings(ctx, settings)
	if err != nil {
		return pkg.LDAPServer{}, fmt.Errorf("failed to load LDAP settings: %w", err)
	}

	for _, setting := range ldapSettings {
		for key, value := range setting {
			switch key {
			case pkg.LDAP_URL:
				ldap.ServerUrl = value
			case pkg.LDAP_PASSWORD:
				ldap.BindPass = value
			case pkg.LDAP_BINDDN:
				ldap.BindDN = value
			case pkg.LDAP_FILTER:
				ldap.UserFilter = value
			case pkg.LDAP_OU:
				ldap.BaseDN = value
			case pkg.LDAP_USERATTR:
				ldap.UserAttr = value
			}
		}
	}
	return ldap, nil
}
