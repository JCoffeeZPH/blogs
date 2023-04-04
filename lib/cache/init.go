package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
	"time"
)

var redisClient *redis.Client

func InitRedis(host, password string, port, db, poolSize, minIdleConns, maxRetries int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
		MaxRetries:   maxRetries,
	})

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stats := client.PoolStats()
	recordRedisPoolStats(stats)
	recordRedisOptions(*(client.Options()))
	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("init redis failed, err: %+v", err))
	}
	redisClient = client
}

func recordRedisPoolStats(stats *redis.PoolStats) {
	if stats != nil {
		logx.Infof("Hits: %d, Misses: %d, Timeout: %d, TotalConns: %d, IdleConns: %d, StaleConns: %d,",
			stats.Hits, stats.Misses, stats.Timeouts, stats.TotalConns, stats.IdleConns, stats.StaleConns)
	}
}

func recordRedisOptions(option redis.Options) {
	optionMap := make(map[string]interface{}, 0)
	t := reflect.TypeOf(option)
	if t.Kind() != reflect.Struct {
		logx.Error("check type err not struct")
		return
	}
	for i := 0; i < t.NumField(); i++ {
		v := reflect.ValueOf(option)
		value := v.FieldByName(t.Field(i).Name)
		if value.Kind() == reflect.Func && !value.IsNil() {
			optionMap[t.Field(i).Name] = "default func"
		} else if value.Kind() == reflect.Int || value.Kind() == reflect.Int64 {
			optionMap[t.Field(i).Name] = value.Int()
		} else if value.Kind() == reflect.Bool {
			optionMap[t.Field(i).Name] = value.Bool()
		} else if value.Kind() == reflect.Func {
			optionMap[t.Field(i).Name] = "default func"
		} else if value.Kind() == reflect.Ptr {
			optionMap[t.Field(i).Name] = "default"
		} else {
			optionMap[t.Field(i).Name] = v.FieldByName(t.Field(i).Name).String()
		}
	}
	logx.Info("redis_utils option info")
}

func GetRedisClient() *redis.Client {
	return redisClient
}
