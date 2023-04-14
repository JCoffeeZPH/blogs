package cache

import (
	"blogs/common/errorx"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

func GetCount(key string) (int64, error) {
	res, err := redisClient.Get(key).Result()
	if err != nil && err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return strconv.ParseInt(res, 10, 64)
}

func IncrCount(key string) {
	redisClient.Incr(key)
}

func SetByKey(key string, value interface{}) {
	_, err := redisClient.Set(key, value, 30*24*time.Hour).Result()
	if err != nil {
		logx.Errorf("SetByKey failed, key: %s, err: %+v", key, err)
	}
}

func GetWebsiteConfig(key string) (string, error) {
	res, err := redisClient.Get(key).Result()
	if err != nil && err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return res, nil
}

func SIsMember(key string, md5 string) bool {
	return redisClient.SIsMember(key, md5).Val()
}

func SAdd(key, md5 string) {
	redisClient.SAdd(key, md5)
}

func HIncr(key, province string) {
	redisClient.HIncrBy(key, province, 1)
}

func HGetAll(key string) map[string]string {
	res, err := redisClient.HGetAll(key).Result()
	if err != nil && err == nil {
		return res
	} else if err != nil {
		panic(errorx.CacheError)
	}
	return res
}
