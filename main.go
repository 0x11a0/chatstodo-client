package main

import ()

func main() {
	setEnv()
	
	gcAPI := initGoogleCalendarApi()
	redisStore := newRedisSessionStore()

	defer redisStore.Close()
	initAuth(redisStore)
	server := initServer(redisStore, gcAPI)
	server.run()
}
