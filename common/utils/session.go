package utils

import (
	"blogs/lib/cache"
	sess "blogs/models/session"
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func ParseJwtToken(token string) string {
	t, err := jwt.Parse(token, nil)
	if t == nil {
		logx.Errorf("parseJwtToken, invalid token: %s, err: %+v", token, err)
		return ""
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		return claims["id"].(string)
	} else {
		return ""
	}
}

func GenerateJwtToken(sessionId, secret string) string {
	now := uint32(time.Now().Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          sessionId,
		"create_time": now,
	})
	key := GenerateJwtTokenKey(secret)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		logx.Errorf("GenerateJwtToken failed, err is: %+v", err)
	}
	return tokenStr
}

func GenerateJwtTokenKey(secret string) []byte {
	md5Value := md5.Sum([]byte(secret))
	signingKey := fmt.Sprintf("%x", md5Value) + "5fE56HFERr44l5EqtE6gLOIuy6%TlaeJ"

	return []byte(signingKey)
}

func ParseSession(jwtToken string) *sess.Session {
	sessionId := ParseJwtToken(jwtToken)
	if len(sessionId) == 0 {
		return nil
	}
	session := cache.SessionBySessionIdFromRedis(sessionId)
	return session
}
