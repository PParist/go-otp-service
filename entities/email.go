package entities

type Email struct {
	Receiver string `validate:"required,email"`
}
type OtpRequest struct {
	Otp_Uid string `validate:"required,uuid4"`
	Otp     string `validate:"required"`
}
