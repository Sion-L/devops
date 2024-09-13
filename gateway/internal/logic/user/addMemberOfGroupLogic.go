package user

import (
	"context"
	"fmt"
	"github.com/Sion-L/devops/user/user"
	"net/http"

	"github.com/Sion-L/devops/gateway/internal/svc"
	"github.com/Sion-L/devops/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMemberOfGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMemberOfGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberOfGroupLogic {
	return &AddMemberOfGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMemberOfGroupLogic) AddMemberOfGroup(req *types.AddMemberOfGroupReq) (resp *types.Response, err error) {
	in := &user.AddMemberOfGroupReq{
		Group: req.Group,
	}

	if _, err = l.svcCtx.User.AddMemberGroup(l.ctx, in); err != nil {
		return nil, err
	}

	return &types.Response{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("新增属性组%s成功", in.Group),
	}, nil
}
