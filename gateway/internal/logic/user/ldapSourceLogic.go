package user

import (
	"context"
	"github.com/Sion-L/devops/user/user"
	"github.com/Sion-L/gateway/internal/svc"
	"github.com/Sion-L/gateway/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type LdapSourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLdapSourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LdapSourceLogic {
	return &LdapSourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LdapSourceLogic) LdapSource(req *types.LdapSourceReq) (resp *types.Response, err error) {
	in := &user.LdapSourceReq{
		Host:     req.Host,
		Port:     req.Port,
		Dn:       req.Dn,
		Password: req.Password,
		Ou:       req.Ou,
		Filter:   req.Filter,
		UserAttr: req.UserAttr,
	}
	if _, err = l.svcCtx.User.LdapSource(l.ctx, in); err != nil {
		return nil, err
	}

	return &types.Response{
		Code:    http.StatusOK,
		Message: "同步ldap用户成功",
	}, nil
}
