package main

import ()

func main() {
	setEnv()
	
	//gcAPI := initGoogleCalendarApi()
	redisStore := newRedisSessionStore()

	defer redisStore.Close()
	server := initServer(redisStore)
	server.run()
}
