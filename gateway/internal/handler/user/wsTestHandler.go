package user

import (
	"net/http"

	"github.com/Sion-L/devops/gateway/internal/logic/user"
	"github.com/Sion-L/devops/gateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WsTestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewWsTestLogic(r.Context(), svcCtx)
		resp, err := l.WsTest()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
