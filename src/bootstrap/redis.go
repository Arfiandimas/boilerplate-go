// Package bootstrap
package bootstrap

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"

	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/logger"
)

// RegistryRedisNative initiate redis session
func RegistryRedisNative() redis.Cmdable {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	r := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:          strings.Split(os.Getenv("REDIS_HOST"), ","),
		ReadTimeout:    time.Duration(2 * time.Second),
		WriteTimeout:   time.Duration(2 * time.Second),
		DB:             db,
		PoolSize:       30,
		PoolTimeout:    time.Duration(10) * time.Second,
		MinIdleConns:   5,
		IdleTimeout:    5 * time.Second,
		RouteByLatency: true,
		Password:       os.Getenv("REDIS_PASSWORD"),
	})

	if r == nil {
		logger.Fatal(`redis cannot connect, please check your config or network`,
			logger.SetField("host", os.Getenv("REDIS_HOST")),
			logger.SetField("db", os.Getenv("REDIS_DB")),
			logger.SetField("name", "redis"),
		)
	}

	c := r.Ping()

	if c.Err() != nil {
		logger.Fatal(`redis cannot connect, please check your config or network`,
			logger.SetField("host", os.Getenv("REDIS_HOST")),
			logger.SetField("db", os.Getenv("REDIS_DB")),
			logger.SetField("name", "redis"),
		)
	}

	return r
}
