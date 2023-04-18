package cache

import (
	"blogs/common/errorx"
	"github.com/go-redis/redis"
)

func ZIncr(key, member string) {
	redisClient.ZIncrBy(key, 1, member)
}

func ZScore(key, member string) float64 {
	res, err := redisClient.ZScore(key, member).Result()
	if err != nil && err == redis.Nil {
		return 0
	} else if err != nil {
		panic(errorx.CacheError)
	}
	return res
}

func ZRevRangeWithScores(key string, start, end int64) []redis.Z {
	res, err := redisClient.ZRevRangeWithScores(key, start, end).Result()
	if err != nil && err == redis.Nil {
		return []redis.Z{}
	} else if err != nil {
		panic(errorx.CacheError)
	}

	return res
}
