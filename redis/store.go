package redis

import (
	"github.com/redis/go-redis/v9"
	"context"
	"log"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis(addr string) {
	rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func SaveCallStatus(callID string, status string) {
	err := rdb.HSet(ctx, "call:"+callID, "status", status).Err()
	if err != nil {
		log.Printf("Error saving call status: %v", err)
	}
}