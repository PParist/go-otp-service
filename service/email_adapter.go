package service

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/PParist/go-otp-service/repositories"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
)

type otpEmailService struct {
	repo repositories.RedisRepository
}

func NewOtpEmailService(repo repositories.RedisRepository) OTPEmailService {
	return &otpEmailService{repo: repo}
}

func (s *otpEmailService) CreateMailOTP(email string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%06d", rand.Intn(1000000)) // Generate a 6-digit OTP
	key := uuid.New().String()
	otpExp := 5 * time.Minute
	err := s.repo.Save(key, otp, otpExp)
	if err != nil {
		fmt.Println("error save")
		return "", err
	}

	if err = sendOtpMail(email, otp); err != nil {
		fmt.Println("error send", err)
		return "", err
	}

	return key, nil
}

func (s *otpEmailService) ValidateOTP(key string, otp string) error {
	result, err := s.repo.Get(key)
	if err != nil {
		return err
	}
	if result != otp {
		return err
	}
	return nil
}

func sendOtpMail(toemail string, otp string) error {
	e := email.NewEmail()
	e.From = "Your Service By kafka <parist.st.work@gmail.com>"
	e.To = []string{toemail}
	e.Subject = "Your OTP Code for Verification"
	e.HTML = []byte(fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Your OTP Code</title>
			<style>
				.container {
					font-family: Arial, sans-serif;
					padding: 20px;
					text-align: center;
					background-color: #f7f7f7;
				}
				.otp-code {
					font-size: 24px;
					font-weight: bold;
					color: #333;
				}
				.message {
					margin-top: 20px;
					font-size: 18px;
					color: #555;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<p class="message">Your OTP code is:</p>
				<p class="otp-code">%s</p>
				<p class="message">Please use this code to complete your verification.</p>
			</div>
		</body>
		</html>
	`, otp))

	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "parist.st.work@gmail.com", os.Getenv("AUTH_GMAIL_PASSWORD"), "smtp.gmail.com"))
	if err != nil {
		fmt.Println("err smtp ", err)
		return err
	}
	return nil
}
