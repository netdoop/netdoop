package store

import (
	"fmt"
	"sync"

	"github.com/netdoop/netdoop/utils"

	"github.com/heypkg/storage"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var defaultRedisClient *redis.Client
var defaultRedisClientOnce sync.Once

var defaultRedisSessionStore *redisstore.RedisStore
var defaultRedisSessionStoreOnce sync.Once

func GetRedisClient() *redis.Client {
	defaultRedisClientOnce.Do(func() {
		env := utils.GetEnv()
		addr := fmt.Sprintf("%v:%v",
			env.GetString("redis_host"),
			env.GetString("redis_port"),
		)
		defaultRedisClient = storage.NewRedisClient(addr)
	})
	return defaultRedisClient
}

func GetRedisSessionStore() *redisstore.RedisStore {
	defaultRedisSessionStoreOnce.Do(func() {
		client := GetRedisClient()
		store, err := storage.GetRedisSessionStore(client, "/acs", 1800)
		if err != nil {
			utils.GetLogger().Fatal("failed to create redis session store", zap.Error(err))
		}
		defaultRedisSessionStore = store
	})
	return defaultRedisSessionStore
}
