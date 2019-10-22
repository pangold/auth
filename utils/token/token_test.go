package token

import (
	"fmt"
	"testing"
)

var which = "JWT"

var token = UseMyToken("127.0.0.1:6379")
// var token = UseJwtToken("MySecret")

func TestGenerateToken(t *testing.T) {
	if token == nil {
		t.Errorf("error: cannot get token")
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

func BenchmarkGenerateToken(b *testing.B) {
	token := getToken()
	if token == nil {
		b.Errorf("error: cannot get token")
	}
	uid, uname, expire := "10001", "dora", 10
	for i := 0; i < b.N; i++ {
		token.GenerateToken(uid, uname, expire)
	}
}

func BenchmarkExplainToken(b *testing.B) {
	token := getToken()
	if token == nil {
		b.Errorf("error: cannot get token")
	}
	var uid, uname string
	tokenString := "C4dFHEsLFdrwadZadBUGZo3zSU8XmpFfTCoD2nA9MTXvwlwbAgXqLYKKLgNhLBSF"
	for i := 0; i < b.N; i++ {
		token.ExplainToken(tokenString, &uid, &uname)
	}
}
