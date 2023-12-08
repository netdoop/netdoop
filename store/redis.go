package store

import (
	"context"
	"fmt"
	"sync"

	"github.com/gorilla/sessions"
	"github.com/netdoop/netdoop/utils"
	"github.com/pkg/errors"

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
		defaultRedisClient = NewRedisClient(addr)
	})
	return defaultRedisClient
}

func NewRedisClient(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return client
}

func GetRedisSessionStore() *redisstore.RedisStore {
	defaultRedisSessionStoreOnce.Do(func() {
		client := GetRedisClient()
		store, err := getRedisSessionStore(client, "/acs", 1800)
		if err != nil {
			utils.GetLogger().Fatal("failed to create redis session store", zap.Error(err))
		}
		defaultRedisSessionStore = store
	})
	return defaultRedisSessionStore
}

func getRedisSessionStore(client *redis.Client, path string, maxAge int) (*redisstore.RedisStore, error) {
	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		return nil, errors.Wrap(err, "new redis store")
	}
	store.KeyPrefix("session:")
	store.Options(sessions.Options{
		Path:     path,
		MaxAge:   maxAge,
		HttpOnly: true,
	})
	return store, nil
}
