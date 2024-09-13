package user

import (
	"context"
	"github.com/Sion-L/devops/gateway/internal/svc"
	"github.com/Sion-L/devops/gateway/internal/types"
	"github.com/Sion-L/devops/user/user"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type LdapVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLdapVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LdapVerifyLogic {
	return &LdapVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LdapVerifyLogic) LdapVerify(req *types.LdapVerifyReq) (resp *types.Response, err error) {
	in := &user.LdapVerifyReq{
		Host:     req.Host,
		Port:     req.Port,
		Dn:       req.Dn,
		Password: req.Password,
		Ou:       req.Ou,
		Filter:   req.Filter,
		UserAttr: req.UserAttr,
	}
	l.Logger.Infof("jwt载体信息,userId: %v,roleType: %v", l.ctx.Value("userId"), l.ctx.Value("roleType"))

	if _, err = l.svcCtx.User.LdapVerify(l.ctx, in); err != nil {
		return nil, err
	}

	return &types.Response{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
