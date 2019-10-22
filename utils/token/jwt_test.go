package token

import (
	"fmt"
	"time"
	"testing"
)

var SecretKey = "MySecretKey"

func Sleep(second int) {
	if second > 0 {
		time.Sleep(time.Duration(second) * time.Second)
	}
}

func TestJwtTokenGenerateToken(t *testing.T) {
	jwtToken := UseJwtToken(SecretKey)
	// step 1
	expectedUid, expectedUname, expireTime := "10001", "dora", 1
	tokenString, _ := jwtToken.GenerateToken(expectedUid, expectedUname, expireTime)
	fmt.Printf("generated token: %s\n", tokenString)
	// expire
	Sleep(-2)
	// step 2
	var userId, userName string
	if err := jwtToken.ExplainToken(tokenString, &userId, &userName); err != nil {
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

func BenchmarkJwtTokenGenerateToken(b *testing.B) {
	jwtToken := UseJwtToken(SecretKey)
	uid, uname, expire := "10001", "dora", 10
	for i := 0; i < b.N; i++ {
		jwtToken.GenerateToken(uid, uname, expire)
	}
}

func BenchmarkJwtTokenExplainToken(b *testing.B) {
	jwtToken := UseJwtToken(SecretKey)
	var uid, uname string
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDE5LTEwLTIyVDEyOjQwOjAwLjg5NzA2OSswODowMCIsInVzZXJJZCI6IjEwMDAxIiwidXNlck5hbWUiOiJkb3JhIn0.lxaLWXJ97vqsHPjBBPGZDG38RDyWOoMq8WibYL65StM"
	for i := 0; i < b.N; i++ {
		jwtToken.ExplainToken(tokenString, &uid, &uname)
	}
}
