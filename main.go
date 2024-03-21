package main

import (
	"github.com:"
)

func main() {
	setEnv()

	redisStore := newRedisSessionStore()
	defer redisStore.Close()
	server := initServer(redisStore)
	server.run()
}
