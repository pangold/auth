package system

import "fmt"

type DefaultVerificationCode struct {

}

func NewDefaultVerificationCode() *DefaultVerificationCode {
	return &DefaultVerificationCode{

	}
}

func (vc *DefaultVerificationCode) SendVerificationCode(vcode string) error {
	fmt.Println("Coming soon.")
	return nil
}
