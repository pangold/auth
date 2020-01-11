package system

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWT as default
type DefaultToken struct {
	SecretKey string
}

func NewDefaultToken(secretKey string) *DefaultToken {
	return &DefaultToken{
		SecretKey: secretKey,
	}
}

func (this *DefaultToken) GenerateToken(uid, name, cid string, expire int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cid":   cid,
		"uid":   uid,
		"uname": name,
		"exp":   time.Now().Add(time.Second * time.Duration(expire)).Unix(),
	})
	tokenString, err := token.SignedString([]byte(this.SecretKey))
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return tokenString, nil
}

func (this *DefaultToken) CheckToken(str string, uid, name, cid *string) error {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(this.SecretKey), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		*cid = claims["cid"].(string)
		*uid = claims["uid"].(string)
		*name = claims["uname"].(string)
	} else {
		return errors.New("unauthorized")
	}
	return nil
}

func (t *DefaultToken) ResetToken(token string) error {
	// Nothing to do with JWT
	return nil
}
