package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/briheet/faultTolerant/server"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func main() {

	// New development Logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Error initilising a new zap development logger: %v", err)
	}

	// Loading of port and all
	err = godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading the .env file: %v", zap.Error(err))
	}

	// New redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password
		DB:       0,  // Default db
	})

	// Check if working or not at start
	err = rdb.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		logger.Error("Error setting redis: %v", zap.Error(err))
	}

	// Initilize the server with required things
	server := server.NewServer(os.Getenv("PORT"), logger, rdb)

	mux := http.NewServeMux()
	server.Serve(mux)

	logger.Info("Starting Server", zap.String("addr", server.Addr))

	if err = http.ListenAndServe(":"+server.Addr, mux); err != nil {
		logger.Error("Error listening and serving",
			zap.String("port", server.Addr),
			zap.Error(err))
	}

}
