package redic

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type MyRedis struct {
	conn redis.Conn
}

func UseMyRedis(hostName string) *MyRedis {
	var myRedis MyRedis
	if err := myRedis.Connect(hostName); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return nil
	}
	return &myRedis
}

func (my *MyRedis) Connect(hostName string) error {
	c, err := redis.Dial("tcp", hostName)
	if err != nil {
		fmt.Println("connection failure")
		return err
	}
	my.conn = c
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
