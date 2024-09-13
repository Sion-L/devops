package user

import "github.com/golang-jwt/jwt/v4"

func GetJwtToken(secretKey string, iat, seconds, refreshAfter, userId int64, roleType int64) (string, error) {
	// 角色映射 塞到token里面去
	// jwtToken, err2 := core.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret,
	//		now, accessExpire, res.UserId, res.RoleType)
	roleMap := map[int64]string{1: "admin", 2: "dev"}
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["refreshAfter"] = refreshAfter // token在什么时候刷新
	claims["userId"] = userId
	claims["role"] = roleMap[roleType]
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
