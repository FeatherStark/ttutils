package ttutils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

// JwtParseClaims 传入JWT字符串，解析并返回其中的claims（明文字段内容），如claims['username']获取username的值
// Args: jwtTokenString JWT字符串, jwtSecret JWT密钥
// Returns: jwt.MapClaims, error
func JwtParseClaims(jwtTokenString, jwtSecret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
