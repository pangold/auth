package system

type DefaultVerificationCode struct {

}

func NewDefaultVerificationCode() *DefaultVerificationCode {
	return &DefaultVerificationCode{

	}
}

func (vc *DefaultVerificationCode) SendVerificationCode(vcode string) error {
	return nil
}
