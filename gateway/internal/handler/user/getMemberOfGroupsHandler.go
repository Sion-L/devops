package user

import (
	"net/http"

	"github.com/Sion-L/devops/gateway/internal/logic/user"
	"github.com/Sion-L/devops/gateway/internal/svc"
	"github.com/Sion-L/devops/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMemberOfGroupsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMemberOfGroupsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetMemberOfGroupsLogic(r.Context(), svcCtx)
		resp, err := l.GetMemberOfGroups(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
