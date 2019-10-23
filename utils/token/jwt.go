package token

import (
	"fmt"
	"time"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	SecretKey string
}

func UseJwtToken(secretKey string) Token {
	return JwtToken{SecretKey: secretKey}
}

func (jt JwtToken) GenerateToken(userId, userName string, expire int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userId,
		"userName": userName,
		"exp":      time.Now().Add(time.Second * time.Duration(expire)).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jt.SecretKey))
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return tokenString, nil
}

func (jt JwtToken) ExplainToken(tokenString string, userId, userName *string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jt.SecretKey), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		*userId = claims["userId"].(string)
		*userName = claims["userName"].(string)
	} else {
		return errors.New("unauthorized")
	}
	return nil
}

func (jt JwtToken) ResetToken(token string) error {
	// Nothing to do with JWT
	return nil
}
