package utils

import (
	"testing"
)

var (
	tokenString string
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("10001", "dora", 1)
	if err != nil {
		t.Errorf("failed to generate token")
	}
	tokenString = token
}

func TestExplainToken(t *testing.T) {
	var uid, uname string
	if err := ExplainToken(tokenString, &uid, &uname); err != nil {
		t.Errorf("failed to explain token")
	}
	if uid != "10001" || uname != "dora" {
		t.Errorf("data error")
	}
}

func TestResetToken(t *testing.T) {
	if err := ResetToken(tokenString); err != nil {
		t.Errorf("failed to reset token")
	}
}
