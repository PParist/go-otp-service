package service

type OTPEmailService interface {
	CreateMailOTP(string) (string, error)
	ValidateOTP(string, string) error
}
