package core

import "github.com/golang-jwt/jwt/v4"

func GetJwtToken(secretKey string, iat, seconds, refreshAfter, userId int64, roleType int64) (string, error) {
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
