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

type LdapVerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLdapVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LdapVerifyLogic {
	return &LdapVerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LdapVerifyLogic) LdapVerify(in *user.LdapVerifyReq) (*user.Empty, error) {
	ldap := &pkg.LDAPServer{
		ServerUrl:  fmt.Sprintf("ldap://%s:%d", in.Host, in.Port),
		BaseDN:     in.Ou,
		BindDN:     in.Dn,
		BindPass:   in.Password,
		UserFilter: in.Filter,
		UserAttr:   fmt.Sprintf(`%s`, in.UserAttr),
	}

	if _, err := ldap.Conn(); err != nil {
		return nil, err
	} else {
		err := l.WriteLdapSettings(in)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (l *LdapVerifyLogic) WriteLdapSettings(in *user.LdapVerifyReq) error {
	settings := []model.Settings{
		{
			Name:      pkg.LDAP_URL,
			Value:     fmt.Sprintf("ldap://%s:%d", in.Host, in.Port),
			Category:  "ldap",
			Encrypted: 0,
		},
		{
			Name:      pkg.LDAP_BINDDN,
			Value:     in.Dn,
			Category:  "ldap",
			Encrypted: 0,
		},
		{
			Name:      pkg.LDAP_OU,
			Value:     in.Ou,
			Category:  "ldap",
			Encrypted: 0,
		},
		{
			Name:      pkg.LDAP_PASSWORD,
			Value:     in.Password,
			Category:  "ldap",
			Encrypted: 0,
		},
		{
			Name:      pkg.LDAP_FILTER,
			Value:     in.Filter,
			Category:  "ldap",
			Encrypted: 0,
		},
		{
			Name:      pkg.LDAP_USERATTR,
			Value:     in.UserAttr,
			Category:  "ldap",
			Encrypted: 0,
		},
	}

	for _, setting := range settings {
		// 检查配置是否已存在
		existingSetting, err := l.svcCtx.SettingsModel.FindOneByName(l.ctx, setting.Name)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return fmt.Errorf("failed to check existing ldap setting: %v", err)
		}

		if existingSetting != nil {
			// 更新现有配置
			existingSetting.Value = setting.Value
			existingSetting.Category = setting.Category
			existingSetting.Encrypted = setting.Encrypted
			if err := l.svcCtx.SettingsModel.Update(l.ctx, existingSetting); err != nil {
				return fmt.Errorf("failed to update ldap setting: %v", err)
			}
		} else {
			// 插入新的配置
			if _, err := l.svcCtx.SettingsModel.Insert(l.ctx, &setting); err != nil {
				return fmt.Errorf("failed to insert ldap settings: %v", err)
			}
		}
	}

	return nil
}
