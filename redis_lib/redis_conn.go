package redis_lib

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

// const params
var (
	RedisClient *redis.Pool
	REDIS_HOST  string
	REDIS_PWD   string
	REDIS_DB    int
)

func init() {
	REDIS_HOST = beego.AppConfig.String("redis::conn")
	REDIS_PWD = beego.AppConfig.String("redis::pwd")
	REDIS_DB, _ = beego.AppConfig.Int("redis::db_num")
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis::maxidle", 500),
		MaxActive:   beego.AppConfig.DefaultInt("redis::maxactive", 3000),
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// password
			_, err = c.Do("AUTH", REDIS_PWD)
			if err != nil {
				c.Close()
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}
