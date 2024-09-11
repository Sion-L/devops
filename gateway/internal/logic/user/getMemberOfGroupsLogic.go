package user

import (
	"context"
	"github.com/Sion-L/devops/user/user"
	"net/http"

	"github.com/Sion-L/gateway/internal/svc"
	"github.com/Sion-L/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberOfGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMemberOfGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberOfGroupsLogic {
	return &GetMemberOfGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMemberOfGroupsLogic) GetMemberOfGroups(req *types.GetMemberOfGroupsReq) (resp *types.ResponseWithData, err error) {
	in := &user.GetMemberOfGroupsReq{}
	groups, err1 := l.svcCtx.User.GetMemberGroups(l.ctx, in)
	if err1 != nil {
		return nil, err1
	}
	return &types.ResponseWithData{
		Code:    http.StatusOK,
		Message: "",
		Data:    groups,
	}, nil
}
