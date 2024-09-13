package middleware

import (
	"fmt"
	core "github.com/Sion-L/devops/core/user"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"time"
)

type RefreshJwtMiddleware struct {
	AccessSecret string
	AccessExpire int64
}

func NewRefreshJwtMiddleware(AccessSecret string, AccessExpire int64) *RefreshJwtMiddleware {
	return &RefreshJwtMiddleware{
		AccessSecret: AccessSecret,
		AccessExpire: AccessExpire,
	}
}

func (m *RefreshJwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取token  标头名字自定义 假设X-Auth
		oldToken := r.Header.Get("X-Auth")
		claims := make(jwt.MapClaims)
		tkn, err := jwt.ParseWithClaims(oldToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.AccessSecret), nil
		})

		if err != nil || !tkn.Valid {
			// 如果 token 无效或过期，则继续执行下一个处理器
			next(w, r)
			return
		}

		now := time.Now().Unix()
		jwtToken, err1 := core.GetJwtToken(m.AccessSecret, now, m.AccessExpire, claims["userId"].(int64), claims["role"].(int64))
		if err1 != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("could not create new token: %v", err))
			return
		}
		w.Header().Set("X-Auth", jwtToken)
		httpx.OkJsonCtx(r.Context(), w, fmt.Sprintf("生成了新token: %s", jwtToken))
		next(w, r)
	}
}
