package caching

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

type RedisConnection struct {
	redisClient *redis.Client
}

// Connection config
func NewConnectionRedis() (conn *RedisConnection) {
	rc := redis.NewClient(&redis.Options{
		Addr:     "192.168.192.2:6379",
		Password: "",
		DB:       0,
	})

	// pong, err := rc.Ping().Result()
	// fmt.Println(pong, err)

	conn = &RedisConnection{rc}
	return conn
}

// Set key
func (conn *RedisConnection) SetRedis(key string, val interface{}, timeEx time.Duration) error {
	err := conn.redisClient.Set(key, val, timeEx*time.Second).Err()

	if err != nil {
		return err
	}

	return nil
}

// Get key
func (conn *RedisConnection) GetRedis(key string) (string, error) {
	val, err := conn.redisClient.Get(key).Result()

	if err != nil {
		return "", errors.New("key not found")
	}

	return val, nil
}
