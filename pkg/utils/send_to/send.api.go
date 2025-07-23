package sendto

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type MailRequest struct {
	ToEmail     string `json:"toEmail"`
	MessageBody string `json:"messageBody"`
	Subject     string `json:"subject"`
	Attachment  string `json:"attachment"`
}

func SendEmailToJavaApi(otp string, email string, purpose string) error {
	// URL of the Java Email API
	postURL := "http://localhost:8080/email/send_text"

	// Build request payload
	mailRequest := MailRequest{
		ToEmail:     email,
		MessageBody: "OTP is: " + otp,
		Subject:     "Verify OTP - " + purpose,
		Attachment:  "path/to/email", // You can leave this empty if not used
	}

	// Convert struct to JSON
	requestBody, err := json.Marshal(mailRequest)
	if err != nil {
		return err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Log response status
	log.Println("Response status:", resp.Status)
	return nil
}
