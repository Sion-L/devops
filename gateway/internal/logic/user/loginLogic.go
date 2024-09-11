package user

import (
	"context"
	"github.com/Sion-L/devops/user/user"
	"github.com/Sion-L/gateway/internal/svc"
	"github.com/Sion-L/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	in := &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}

	res, err1 := l.svcCtx.User.Login(l.ctx, in)
	if err1 != nil {
		return nil, err1
	}
	return &types.LoginResp{
		UserId:   res.UserId,
		Username: res.Username,
	}, nil
}
