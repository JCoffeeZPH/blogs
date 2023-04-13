package session

import (
	"net/http"
	"strconv"
)

type Session struct {
	SessionId     string `json:"session_id"`
	UserId        int64  `json:"user_id"`
	Username      string `json:"username"`
	Secret        string `json:"secret"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	Email         string `json:"email"`
	Role          int8   `json:"role"`
	LastLoginTime uint32 `json:"last_login_time"`
	IpAddress     string `json:"ip_address"`
	IpSource      string `json:"ip_source"`
	IsSubscribe   int8   `json:"is_subscribe"`
	LoginType     int8   `json:"login_type"`
	Website       string `json:"website"`
}

func GetInt64(r *http.Request, key string) int64 {
	v := r.Header.Get(key)
	if len(v) == 0 {
		return -1
	}
	val, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return -1
	}
	return val
}
