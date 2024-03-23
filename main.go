package main

import (
	"log"

	"github.com/lucasodra/chatstodo-client/internal"
	"github.com/lucasodra/chatstodo-client/internal/utils"
)

var requiredEnv = []string{
	"CSRF_SECRET",
	"REDIS_ADDR",
	"REDIS_SECRET",
	"GOOGLE_KEY",
	"GOOGLE_SECRET",
	"GOOGLE_CALLBACK_URL",
	"BACKEND_AUTH_GET_JWT_URL",
	"BACKEND_GET_SUMMARY_URL",
	"BACKEND_GET_ALL_PLATFORM_URL",
	"BACKEND_ADD_PLATFORM_URL",
	"BACKEND_REMOVE_PLATFORM_URL",
	"BACKEND_GET_ALL_GROUPS_URL",
	"BACKEND_DELETE_GROUP_URL",
}

func main() {
	envMap := utils.SetEnv()
	for _, envVar := range requiredEnv {
		if !envMap[envVar] {
			log.Fatalln("Required env variable " + envVar + " not found. Exiting.")
		}

	}

	redisStore := utils.NewRedisSessionStore()
	defer redisStore.Close()
	server := internal.InitServer(redisStore)
	server.Run()
}
