package redis

import (
	"context"
	"fmt"
	"strconv"

	"awesomeProject/helper"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func InitRedis() error {

	redisHost := helper.GetENV("REDIS_HOST", "127.0.0.1")
	redisPort := helper.GetENV("REDIS_PORT", "6379")
	redisPassword := helper.GetENV("REDIS_PASSWORD", "")
	redisDB := helper.GetENV("REDIS_DB", "0")

	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	db, err := strconv.Atoi(redisDB)
	if err != nil {
		return err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisPassword,
		DB:       db,
	})

	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}

	Rdb = rdb
	fmt.Println("Successfully connected to Redis!")

	return nil
}
