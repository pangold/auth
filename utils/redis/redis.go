package redis

import (
	"github.com/garyburd/redigo/redis"
)

type MyRedis struct {
	conn *Redis.Conn
}

func (my *MyRedis) Connect(hostName string) error {
	c, err := redis.Dail("tcp", hostName)
	if err != nil {
		fmt.Println("connection failure")
		return err
	}
	my.conn = &c
	return nil
}

func (my *MyRedis) Close() error {
	return my.conn.Close()
}

func (my *MyRedis) Set(key string, value interface{}) error {
	return nil
}

func (my *MyRedis) Get(key string, value *interface{}) error {
	return nil
}
