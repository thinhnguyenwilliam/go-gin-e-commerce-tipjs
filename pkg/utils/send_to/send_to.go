package sendto

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/thinhcompany/ecommerce-ver-2/global"
	"go.uber.org/zap"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

// Replace these constants with your actual SMTP settings
const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = 587
	SMTPUsername = "thinhproee@gmail.com"
	SMTPPassword = "rmakdncnfplgdrfr" // App Password, not account password
)

// Builds MIME-formatted email content
func BuildMessage(mail Mail) string {
	msg := "MIME-Version: 1.0\r\n"
	msg += "Content-Type: text/html; charset=\"UTF-8\"\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += "\r\n" + mail.Body + "\r\n"
	return msg
}

// Main function to send OTP email using a template
func SendTemplateOtp(to []string, from string, templateName string, data map[string]any) error {
	htmlBody, err := getEmailTemplate(templateName, data)
	if err != nil {
		return err
	}
	return Send(to, from, htmlBody)
}

// Load HTML template and render with data
func getEmailTemplate(templateName string, data map[string]any) (string, error) {
	htmlBuffer := new(bytes.Buffer)
	tmpl, err := template.ParseFiles("template-email/" + templateName)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(htmlBuffer, data)
	if err != nil {
		return "", err
	}
	return htmlBuffer.String(), nil
}

// Sends the actual email
func Send(to []string, from string, htmlBody string) error {
	contentEmail := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "Test Sender",
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlBody,
	}

	messageEmail := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", SMTPHost, SMTPPort),
		auth,
		from,
		to,
		[]byte(messageEmail),
	)

	if err != nil {
		global.AppLogger.Error("Email send failed", zap.Error(err))
		return err
	}
	return nil
}

// Sends a plain OTP email
func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "Test Sender",
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is <b>%s</b>. Please enter it to verify your account.", otp),
	}

	messageEmail := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", SMTPHost, SMTPPort),
		auth,
		from,
		to,
		[]byte(messageEmail),
	)

	if err != nil {
		//fmt.Println("Email send failed:", err)
		global.AppLogger.Error("Email send failed", zap.Error(err))
		return err
	}

	return nil
}
