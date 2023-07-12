package signup

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gopkg.in/mail.v2"
)

var Sender EmailSender

// EmailSender store information of the sender who send verify email
type EmailSender struct {
	Mail       *mail.Message
	SmtpServer string
	SmtpPort   int
	From       string
	Password   string
}

// VerifyEmail have the data required to verify an email
type VerifyEmail struct {
	Email      string
	SecretCode string
}

// InitFromDotenv is a constructor of EmailSender
func (e *EmailSender) InitFromDotenv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	e.From = os.Getenv("EMAIL_SENDER")
	e.Password = os.Getenv("EMAIL_PASS")
	e.SmtpServer = os.Getenv("EMAIL_SMTP")
	e.SmtpPort, err = strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}
	return nil
}

// SendEmailVerify implements send verify email from EmailSender
func (e *EmailSender) SendEmailVerify(v VerifyEmail) error {
	e.Mail.SetHeader("From", e.From)
	e.Mail.SetHeader("To", v.Email)
	e.Mail.SetHeader("Subject", "Verification Email")
	verifyLink := fmt.Sprintf("http://localhost:8081/v1/verify?verification_code=%s", v.SecretCode)
	msg := fmt.Sprintf(`Your Verification link: <a href="%s">click here</a>`, verifyLink)
	e.Mail.SetBody("text/html", msg)

	d := mail.NewDialer(e.SmtpServer, e.SmtpPort, e.From, e.Password)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	if err := d.DialAndSend(e.Mail); err != nil {
		return err
	}
	return nil
}

// ValidataVerifyCode check the secret code is existed in Redis record or not
func ValidateVerifyCode(ctx context.Context, r *redis.Client, code string) (string, bool) {
	res := r.Get(ctx, code)
	if res.Err() == redis.Nil {
		return "", false
	}
	return res.Val(), true
}
