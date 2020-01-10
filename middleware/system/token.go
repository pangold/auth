package system

// JWT as default
type DefaultToken struct {
	SecretKey string
}

func NewDefaultToken(secretKey string) *DefaultToken {
	return &DefaultToken{
		SecretKey: secretKey,
	}
}

func (t *DefaultToken) GenerateToken(id, name, cid string) (string, error) {
	return "", nil
}

func (t *DefaultToken) TokenVerification(token string, id, name, cid *string) error {
	return nil
}

func (t *DefaultToken) ResetToken(token string) error {
	return nil
}
