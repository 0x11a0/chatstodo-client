package main

import ()

func main() {
	setEnv()

	redisStore := newRedisSessionStore()
	defer redisStore.Close()
	initAuth(redisStore)
	server := initServer(redisStore)
	server.run()
}
