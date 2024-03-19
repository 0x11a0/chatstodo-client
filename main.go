package main

import ()

func main() {
	setEnv()
	
	redisStore := newRedisSessionStore()
	defer redisStore.Close()
	server := initServer(redisStore)
	server.run()
}
