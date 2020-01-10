package system

type DefaultEmail struct {
	Name string
	Password string
	Host string
	Port int
}

func NewDefaultEmailService(name, pwd, host string, port int) *DefaultEmail {
	return &DefaultEmail{
		Name: name,
		Password: pwd,
		Host: host,
		Port: port,
	}
}

func (vc *DefaultEmail) SendVerificationCode(vcode string) error {
	return nil
}

func (vc *DefaultEmail) SendActivationEmail(url string) error {
	return nil
}

func (vc *DefaultEmail) SendResetPasswordEmail(url string) error {
	return nil
}