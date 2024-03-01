package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
	"gopkg.in/boj/redistore.v1"
)

type RedisSessionStore struct {
	store *redistore.RediStore
}

func newRedisSessionStore() *RedisSessionStore {
	store, err := redistore.NewRediStore(10, "tcp", ":6379",
		os.Getenv("REDIS_PASSWORD"), []byte(os.Getenv("REDIS_SECRET")))
	if err != nil {
		log.Println("db.go - newRedisStore()")
		log.Fatal(err)
	}
	store.SetMaxAge(86400 * 30)
	store.Options.SameSite = http.SameSiteDefaultMode
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = os.Getenv("IS_PROD") == "true"
	log.Println("RedisStorage connected successfully")

	return &RedisSessionStore{
		store: store,
	}
}

func (redisSessionStore *RedisSessionStore) close() {
	redisSessionStore.store.Close()
}

type RedisStorage struct {
	client *redis.Client
}

func newRedisStorage() *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return &RedisStorage{
		client: client,
	}
}

func (redisClient *RedisStorage) set(key string, value string) error {
	return redisClient.client.Set(context.Background(), key, value, 0).Err()
}

func (redisClient *RedisStorage) get(key string) (string, error) {
	return redisClient.client.Get(context.Background(), key).Result()
}

func (redisClient *RedisStorage) close() {
	redisClient.client.Close()
}
