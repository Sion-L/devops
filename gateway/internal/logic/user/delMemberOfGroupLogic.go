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

type DelMemberOfGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelMemberOfGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMemberOfGroupLogic {
	return &DelMemberOfGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelMemberOfGroupLogic) DelMemberOfGroup(req *types.DelMemberOfGroupReq) (resp *types.Response, err error) {
	in := &user.DelMemberOfGroupReq{
		Group: req.Group,
	}

	if _, err = l.svcCtx.User.DelMemberGroup(l.ctx, in); err != nil {
		return nil, err
	}

	return &types.Response{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("成功删除属性组%s", in.Group),
	}, nil
}
