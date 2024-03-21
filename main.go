package main

import (
	"github.com/lucasodra/chatstodo-client/internal/utils"
	"github.com/lucasodra/chatstodo-client/internal"
)

func main() {
	utils.SetEnv()

	redisStore := utils.NewRedisSessionStore()
	defer redisStore.Close()
	server := internal.InitServer(redisStore)
	server.Run()
}
