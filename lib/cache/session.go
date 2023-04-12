package cache

import (
	"blogs/common/constants"
	sess "blogs/models/session"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
)

func sessionRedisKey(sessionId string) string {
	return fmt.Sprintf("blob.user_session.%s", sessionId)
}

func SessionBySessionIdFromRedis(sessionId string) *sess.Session {
	session, err := redisClient.Get(sessionRedisKey(sessionId)).Result()
	if err == redis.Nil {
		logx.Debugf("SessionBySessionIdFromRedis not found session, sessionId is %s", sessionId)
		return nil
	}
	s := &sess.Session{}
	err = json.Unmarshal([]byte(session), s)
	if err != nil {
		logx.Errorf("SessionBySessionIdFromRedis failed, err is %+v", err)
		return nil
	}
	return s
}

func SaveSessionIntoRedis(session *sess.Session) {
	key := sessionRedisKey(session.SessionId)
	sessionByte, _ := json.Marshal(session)
	redisClient.Set(key, sessionByte, constants.SessionDefaultTimeOut)
}
