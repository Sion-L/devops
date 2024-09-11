package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/Sion-L/devops/user/model"

	"go.uber.org/zap"

	pkg "github.com/Sion-L/devops/pkg/user"
	"github.com/Sion-L/devops/user/internal/svc"
	"github.com/Sion-L/devops/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LdapSourceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLdapSourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LdapSourceLogic {
	return &LdapSourceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LdapSourceLogic) LdapSource(in *user.LdapSourceReq) (*user.Empty, error) {
	if err := l.SyncLdap(in); err != nil {
		return &user.Empty{}, err
	}
	return &user.Empty{}, nil
}

func (l *LdapSourceLogic) SyncLdap(in *user.LdapSourceReq) error {
	ldap := &pkg.LDAPServer{
		ServerUrl:  fmt.Sprintf("ldap://%s:%d", in.Host, in.Port),
		BaseDN:     in.Ou,
		BindDN:     in.Dn,
		BindPass:   in.Password,
		UserFilter: in.Filter,
		UserAttr:   in.UserAttr,
	}

	usersAttr, err := ldap.SearchLDAPUsers()
	if err != nil {
		return err
	}

	for _, userAttr := range usersAttr {
		existingUser, err := l.svcCtx.AuthUsersModel.FindOneByUser(l.ctx, userAttr.Username)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return err
		}

		l.Logger.Infof("某些用户没有mobile字段: %v", userAttr)

		userInfo := &model.AuthUsers{
			UserId:   pkg.GenerateUserId(1),
			Username: userAttr.Username,
			Password: "",
			NickName: userAttr.NickName,
			RoleName: "管理员",
			Source:   "ldap",
			RoleType: 1,
			Email:    userAttr.Email,
			Mobile:   userAttr.Mobile, // 特殊场景下某些用户ldap中没有这个字段
		}

		if existingUser != nil {
			l.Logger.Info("查出来的用户: %v", existingUser)
			userInfo.UserId = existingUser.UserId // Keep the existing user ID
			err = l.svcCtx.AuthUsersModel.Update(l.ctx, userInfo)
			if err != nil {
				return err
			}
			l.Logger.Info("[mysql] Updated existing LDAP user information: ", zap.Any("user", userInfo))
		} else {
			_, err := l.svcCtx.AuthUsersModel.Insert(l.ctx, userInfo)
			if err != nil {
				return err
			}
			l.Logger.Info("[mysql] Inserted new LDAP user information: ", zap.Any("user", userInfo))
		}
	}

	return nil
}
