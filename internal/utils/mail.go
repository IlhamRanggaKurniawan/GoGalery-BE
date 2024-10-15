package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/smtp"
	"os"
)

func GenerateOtp() (string, error) {
	otp := ""
	digits := "0123456789"

	for i := 0; i < 6; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp += string(digits[num.Int64()])
	}

	return otp, nil
}

func SendEmailOTP(to string, otp string) error {
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"
    senderEmail := os.Getenv("EMAIL_SENDER")
    senderPassword := os.Getenv("EMAIL_APP_PASSWORD")


    subject := "Subject: OTP for Password Reset\n"

    body := fmt.Sprintf(`
    <html>
    <body>
      <h1>Password Reset OTP</h1>
      <p>Here is your OTP to reset your password:</p>
      <h2>%s</h2>
      <p>This OTP is valid for 5 minutes. Please do not share it with anyone.</p>
    </body>
    </html>
    `, otp)

    message := []byte(subject + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + body)

    auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{to}, message)
    if err != nil {
        return fmt.Errorf("failed to send email: %w", err)
    }

    return nil
}
