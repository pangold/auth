package verifier

import (
	"fmt"
	"testing"
)

func TestSendEmail(t *testing.T) {
	body := "Hi King, <br><br>This is a test. <br><br>From your most sincere friend"
	e := UseGomail("smtp.163.com", 485, "pangold@163.com", "******")
	if err := e.SendEmail("Hello", body, "pangold@163.com"); err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("bingo")
}
