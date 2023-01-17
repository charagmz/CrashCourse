package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/charagmz/CrashCourse/entity"
	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *entity.Post) {
	// Create redis client
	client := cache.getClient()

	// Serialize the post that we receive
	json, err := json.Marshal(value)
	if err != nil {
		log.Printf("Error serializing the post %v", err.Error())
	} else {
		// Save key and value in redis
		client.Set(key, json, cache.expires*time.Second)
	}

}

func (cache *redisCache) Get(key string) *entity.Post {
	// Create redis client
	client := cache.getClient()
	//fmt.Printf("id key %v", key)

	// Get the post from redis
	val, err := client.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("Doesn't exists the post %v", err.Error())
		} else {
			log.Printf("Error getting the post %v", err.Error())
		}
		return nil
	}

	// unmarshaling post
	post := entity.Post{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		log.Printf("Error unmarshaling the post %v", err.Error())
		return nil
	}

	return &post
}
