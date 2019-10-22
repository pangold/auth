package token

import (
	"fmt"
	"testing"
)

var host = "127.0.0.1:6379"

func TestMyTokenGenerateToken(t *testing.T) {
	token := UseMyToken(host)
	if token == nil {
		t.Errorf("error: cannot get my token")
	}
	// step 1
	expectedUid, expectedUname, expireTime := "10001", "dora", 1
	tokenString, _ := token.GenerateToken(expectedUid, expectedUname, expireTime)
	fmt.Printf("generated token: %s\n", tokenString)
	// expire
	Sleep(-2)
	// step 2
	var userId, userName string
	if err := token.ExplainToken(tokenString, &userId, &userName); err != nil {
		t.Errorf("error: %s", err.Error())
	}
	// compare
	if userId != expectedUid {
		t.Errorf("userId [%s] expected [%s]", userId, expectedUid)
	}
	if userName != expectedUname {
		t.Errorf("userName [%s] expected [%s]", userName, expectedUname)
	}
}

func BenchmarkMyTokenGenerateToken(b *testing.B) {
	token := UseMyToken(host)
	if token == nil {
		b.Errorf("error: cannot get my token")
	}
	uid, uname, expire := "10001", "dora", 10
	for i := 0; i < b.N; i++ {
		token.GenerateToken(uid, uname, expire)
	}
}

func BenchmarkMyTokenExplainToken(b *testing.B) {
	token := UseMyToken(host)
	if token == nil {
		b.Errorf("error: cannot get my token")
	}
	var uid, uname string
	tokenString := "s9P77mlbx4HwXRbsbAqFICRbBNrusx0de367dJ7EPzANnifSzqo5cPDcdlQiIAC9"
	for i := 0; i < b.N; i++ {
		token.ExplainToken(tokenString, &uid, &uname)
	}
}
