package middleware

// Default is using JWT
type Token interface {
	// cid could be: pc, android, ios, web
	GenerateToken(id, name, cid string) (string, error)
	TokenVerification(token string, id, name, cid *string) error
	ResetToken(token string) error
}

// To customize the content of Verification Code
// Manage your 3rd SMS Service / Email Service yourself
type VerificationCode interface {
	SendVerificationCode(vcode string) error   // less than 70 characters if SMS
}

// To customize the content of your Email
// How do you want to show your users?
// Customize it, or use default.
// And to manage your Email service
// Default email service: gomail
type Email interface {
	SendActivationEmail(url string) error
	SendResetPasswordEmail(url string) error
}

type ThirdParty interface {

}

// If needs email activation, do it here,
// Others, Doesn't need to do anything
type AfterSignedUp interface {

}

type AfterSignedIn interface {

}

type AfterSignedOut interface {

}

// Extend for registration successfully,
// such as: welcome message(via email / mobile SMS / IM...), new user guide, award, recommend awards
// Default: write into database, otherwise do nothing.
type AfterRegistration interface {
	Done()
}