package token


type Token interface {
	GenerateToken(userId, userName string, expire int) (string, error)
	ExplainToken(token string, userId, userName *string) error
	ResetToken(token string) error
}
