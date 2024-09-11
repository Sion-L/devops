package user

import (
	"net/http"

	"github.com/Sion-L/gateway/internal/logic/user"
	"github.com/Sion-L/gateway/internal/svc"
	"github.com/Sion-L/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddUserToMemberOfGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserToMemberOfGroupReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewAddUserToMemberOfGroupLogic(r.Context(), svcCtx)
		resp, err := l.AddUserToMemberOfGroup(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
