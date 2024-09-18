package middleware

import (
	"fmt"
	"github.com/Sion-L/devops/core"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"time"
)

type RefreshJwtMiddleware struct {
	AccessSecret         string
	AccessExpire         int64
	TokenDisableDuration int64
}

func NewRefreshJwtMiddleware(AccessSecret string, AccessExpire int64, TokenDisableDuration int64) *RefreshJwtMiddleware {
	return &RefreshJwtMiddleware{
		AccessSecret:         AccessSecret,
		AccessExpire:         AccessExpire,
		TokenDisableDuration: TokenDisableDuration,
	}
}

func (m *RefreshJwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取token 标头名字自定义 假设X-Auth
		oldToken := r.Header.Get("X-Auth")
		claims := make(jwt.MapClaims)

		// 解析并验证 JWT
		tkn, err := jwt.ParseWithClaims(oldToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.AccessSecret), nil
		})

		if err != nil || !tkn.Valid {
			// 如果 token 无效或过期，则继续执行下一个处理器
			next(w, r)
			return
		}

		// 获取当前时间和 token 的创建时间
		now := time.Now().Unix()
		iat := int64(claims["iat"].(float64)) // 获取 token 的 iat (issued at) 字段

		// 检查 token 是否超过禁用时长 超过是三天则token被禁用
		if now-iat > m.TokenDisableDuration {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("token has been disabled due to exceeding the allowed age"))
			return
		}

		// token刷新时机: 当前时间大于刷新时间则刷新token
		refreshAfter := int64(claims["refreshAfter"].(float64))
		if now > refreshAfter {
			newRefreshAfter := now + m.AccessExpire/2
			jwtToken, err1 := core.GetJwtToken(m.AccessSecret, now, m.AccessExpire, newRefreshAfter, claims["userId"].(int64), claims["role"].(int64))
			if err1 != nil {
				httpx.ErrorCtx(r.Context(), w, fmt.Errorf("could not create new token: %v", err))
				return
			}
			w.Header().Set("X-Auth", jwtToken)
			httpx.OkJsonCtx(r.Context(), w, fmt.Sprintf("生成了新token: %s", jwtToken))
		}

		// 继续执行下一个处理器
		next(w, r)
	}
}
