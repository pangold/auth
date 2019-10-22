package token

import (
	"fmt"
	"errors"
	"strings"

	"github.com/gomodule/redigo/redis"
)

type MyToken struct {
	conn redis.Conn
}

func UseMyTokenWithError(hostName string) (Token, error) {
	fmt.Println("use my token")
	c, err := redis.Dial("tcp", hostName)
	if err != nil {
		return nil, err
	}
	// FIXME: conn.Close()
	return MyToken{conn: c}, nil
}

func UseMyToken(hostName string) Token {
	token, err := UseMyTokenWithError(hostName)
	if err != nil {
		return nil
	}
	return token
}

func (mt MyToken) GenerateToken(userId, userName string, expire int) (string, error) {
	key := GenerateRandomString(64) // key is token
	value := userId + "," + userName
	if _, err := mt.conn.Do("SET", key, value); err != nil {
		return "", err
	}
	// set expire time
	if _, err := mt.conn.Do("EXPIRE", key, expire); err != nil {
		mt.conn.Do("DEL", key)
		return "", err
	}
	return key, nil
}

func (mt MyToken) ExplainToken(token string, userId, userName *string) error {
	v, err := redis.String(mt.conn.Do("GET", token))
	if err != nil {
		return err
	}
	kv := strings.Split(v, ",")
	if len(kv) != 2 {
		return errors.New("invalid value")
	}
	*userId, *userName = kv[0], kv[1]
	return nil
}

func (mt MyToken) ResetToken(token string) error {
	if _, err := mt.conn.Do("DEL", token); err != nil {
		return err
	}
	return nil
}
