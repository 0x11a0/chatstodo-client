package main

import ()

func main() {
	setEnv()

	redisStore := newRedisSessionStore()
	server := initServer(redisStore)
	server.run()
}
