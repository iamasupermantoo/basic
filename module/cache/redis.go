package cache

import (
	"basic/config"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var Rds *redis.Pool

// InitRedis 初始化Redis
func InitRedis(conf *config.Redis) {
	Rds = &redis.Pool{
		MaxIdle:     conf.MaxIdleConn,
		MaxActive:   conf.MaxOpenConn,
		IdleTimeout: time.Duration(conf.ConnectTimeout) * time.Second,
		Wait:        conf.Wait,
		Dial: func() (redis.Conn, error) {
			host := fmt.Sprintf("%s:%d", conf.Server, conf.Port)
			conn, err := redis.Dial(
				conf.Network,
				host,
				redis.DialPassword(conf.Pass),
				redis.DialDatabase(conf.DbName),
				redis.DialConnectTimeout(time.Duration(conf.ConnectTimeout)*time.Second),
				redis.DialReadTimeout(time.Duration(conf.ReadTimeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(conf.WriteTimeout)*time.Second),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}
