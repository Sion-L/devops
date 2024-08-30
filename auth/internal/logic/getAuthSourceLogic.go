package logic

import (
	"context"

	"github.com/Sion-L/auth/auth"
	"github.com/Sion-L/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthSourceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuthSourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthSourceLogic {
	return &GetAuthSourceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuthSourceLogic) GetAuthSource(in *auth.AuthSource) (*auth.TypeResp, error) {
	// todo: add your logic here and delete this line

	return &auth.TypeResp{}, nil
}
