package user

import (
	"context"
	"net/http"

	"github.com/Sion-L/devops/gateway/internal/svc"
	"github.com/Sion-L/devops/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WsTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsTestLogic {
	return &WsTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsTestLogic) WsTest() (resp *types.Response, err error) {
	return &types.Response{
		Code:    http.StatusOK,
		Message: "websocket测试",
	}, nil
}
