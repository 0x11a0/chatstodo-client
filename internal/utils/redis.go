package utils

import (
	"gopkg.in/boj/redistore.v1"
	"log"
	"net/http"
	"os"
)


func NewRedisSessionStore() *redistore.RediStore {
	store, err := redistore.NewRediStore(10, "tcp", os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PASSWORD"), []byte(os.Getenv("REDIS_SECRET")))
	if err != nil {
		log.Println("redis.go - newRedisSessionStore()")
		log.Fatal(err)
	}
	store.SetMaxAge(86400 * 30)
	//store.Options.SameSite = http.SameSiteDefaultMode
	store.Options.SameSite = http.SameSiteStrictMode
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = os.Getenv("IS_PROD") == "true"
	log.Println("RedisSessionStore connected successfully")
	return store
}
