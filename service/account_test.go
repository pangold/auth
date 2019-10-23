package service

import (
	"testing"
)

func TestRequirePhoneCode(t *testing.T) {
	if err := RequirePhoneCode("18828883888"); err != nil {
		t.Errorf("failed to require phone verification code")
	}
}
