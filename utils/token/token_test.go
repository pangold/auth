package token

import (
	"fmt"
	"time"
	"testing"
)

var (
	token Token
	tokenString string
)

func init() {
	token = UseJwtToken("MySecret")
	//token = UseMyToken("127.0.0.1:6379")
	if token == nil {
		panic("invalid token object")
	}
}

func TestGenerateToken(t *testing.T) {
	uid, uname, expire := "10001", "dora", 1
	ts, err := token.GenerateToken(uid, uname, expire)
	if err != nil {
		t.Errorf("failed to generate token")
		return
	}
	tokenString = ts
	fmt.Printf("generated token: %s\n", ts)
}

func TestExplainToken1(t *testing.T) {
	var userId, userName string
	expectedUid, expectedUname := "10001", "dora"
	if err := token.ExplainToken(tokenString, &userId, &userName); err != nil {
		t.Errorf("parse error: %s", err.Error())
		return
	}
	if userId != expectedUid {
		t.Errorf("uid [%s] didn't match expected uid [%s]", userId, expectedUid)
	}
	if userName != expectedUname {
		t.Errorf("uname [%s] didn't match expected uname [%s]", userName, expectedUname)
	}
}

func TestExplainToken2(t *testing.T) {
	time.Sleep(time.Duration(2) * time.Second)
	var userId, userName string
	expectedUid, expectedUname := "10001", "dora"
	token.ExplainToken(tokenString, &userId, &userName)
	// expire time only 1s, it has sleep 2s in front, so...
	if userId == expectedUid {
		t.Errorf("uid [%s] shouldn't match expected uid [%s]", userId, expectedUid)
	}
	if userName == expectedUname {
		t.Errorf("uname [%s] shouldn't match expected uname [%s]", userName, expectedUname)
	}
}

func BenchmarkJwtTokenGenerateToken(b *testing.B) {
	uid, uname, expire := "10001", "dora", 60 * 60 // an hour
	for i := 0; i < b.N; i++ {
		tokenString, _ = token.GenerateToken(uid, uname, expire)
	}
}

func BenchmarkJwtTokenExplainToken(b *testing.B) {
	var uid, uname string
	for i := 0; i < b.N; i++ {
		token.ExplainToken(tokenString, &uid, &uname)
	}
}
