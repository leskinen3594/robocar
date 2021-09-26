package caching

import (
	"time"

	"github.com/go-redis/redis"
)

// Connection config
func ConnectRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "192.168.192.2:6379",
		Password: "",
		DB:       0,
	})

	// pong, err := redisClient.Ping().Result()
	// fmt.Println(pong, err)

	return redisClient
}

// Set key
func SetRedis(rc *redis.Client, key string, val interface{}, timeEx time.Duration) error {
	err := rc.Set(key, val, timeEx*time.Second).Err()

	if err != nil {
		return err
	}

	return nil
}

// Get key
func GetRedis(rc *redis.Client, key string) (string, error) {
	val, err := rc.Get(key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}
