package model

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)
var RedisConn *redis.Pool

type Redis struct {

}

func init() {
	RedisConn = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 200,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
	return
}

//存入
func (r *Redis)Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value, "EX",time)
	//conn.Flush();
	if err != nil {
		return err
	}
	return err
}


//判断是否存在
func (r *Redis) Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

//取数据
func(r *Redis) Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

//删除
func (r *Redis) Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("DEL", key))
}

