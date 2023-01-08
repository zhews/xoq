package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/go-redis/redis/v9"
	"log"
	"os"
)

var (
	policyFileName string
	redisHost      string
	redisUsername  string
	redisPassword  string
)

func init() {
	flag.StringVar(&policyFileName, "policy", "policy.json", "The policy in form of a JSON that should be imported into the Redis Q-Table.")
	flag.StringVar(&redisHost, "redis-host", "127.0.0.1:6379", "The host of the Redis server you want to import the policy to.")
	flag.StringVar(&redisUsername, "redis-username", "default", "The username to use while connecting to the Redis server.")
	flag.StringVar(&redisPassword, "redis-password", "", "The password to use while connecting to the Redis server.")
	flag.Parse()
}

func main() {
	log.Println("Starting to import", policyFileName, "to", redisHost, "...")
	policyFile, err := os.ReadFile(policyFileName)
	if err != nil {
		log.Fatalln("Could not open policy file!", err)
	}
	var policy map[string]float64
	if err := json.Unmarshal(policyFile, &policy); err != nil {
		log.Fatalln("Could not parse policy JSON!", err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Username: redisUsername,
		Password: redisPassword,
	})
	for key, value := range policy {
		status := redisClient.Set(context.Background(), key, value, 0)
		log.Println(status.Result())
		if status.Err() != nil {
			log.Fatalln("Could not set key", key, "with value", value, "because of an error.", err)
		}
	}
}
