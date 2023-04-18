package middleware

import (
	"blogs/common/constants"
	"blogs/common/errorx"
	"blogs/common/utils"
	"blogs/lib/cache"
	sess "blogs/models/session"
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

type AuthTokenMiddleware struct {
}

func NewAuthTokenMiddleware() *AuthTokenMiddleware {
	return &AuthTokenMiddleware{}
}

func (m *AuthTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.String(), "users/login") || strings.Contains(r.URL.String(), "/report") {
			next(w, r)
		}

		jwtToken := getJwtToken(r)
		if len(jwtToken) == 0 {
			panic(errorx.Unauthorized)
		}
		s := utils.ParseSession(jwtToken)
		if s == nil {
			panic(errorx.Unauthorized)
		}
		auth(jwtToken, s, r)
		next(w, r)
	}
}

func getJwtToken(r *http.Request) string {
	header := r.Header.Get("Authorization")
	splits := strings.Split(header, " ")
	if len(splits) == 2 && strings.ToLower(splits[0]) == "bearer" {
		return splits[1]
	} else {
		return ""
	}
}

func auth(jwtToken string, s *sess.Session, r *http.Request) {
	if !verifyJwtToken(jwtToken, s.Secret) {
		panic(errorx.Unauthorized)
	}

	s.LastLoginTime = uint32(time.Now().Unix())
	go cache.SaveSessionIntoRedis(s)
	// todo 投递消息队列更新数据库user_info_tab的last_login_time

	setHeader(r, s)
}

func verifyJwtToken(jwtToken, secret string) bool {
	md5Value := md5.Sum([]byte(secret))
	md5 := fmt.Sprintf("%x", md5Value)

	signingKey := md5 + "5fE56HFERr44l5EqtE6gLOIuy6%TlaeJ"

	_, e := jwt.Parse(jwtToken, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(signingKey), nil
	})

	if e != nil {
		return false
	} else {
		return true
	}
}

func setHeader(r *http.Request, s *sess.Session) {
	r.Header.Set(constants.UserId, fmt.Sprintf("%d", s.UserId))
}
