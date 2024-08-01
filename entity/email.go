package entity

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string, attachmentPath string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpEmail := os.Getenv("SMTP_EMAIL")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	fmt.Println("SMTP_SERVER:", smtpServer)
	fmt.Println("SMTP_PORT:", smtpPort)
	fmt.Println("SMTP_EMAIL:", smtpEmail)
	fmt.Println("SMTP_PASSWORD:", smtpPassword)

	if smtpServer == "" || smtpPort == "" || smtpEmail == "" || smtpPassword == "" {
		fmt.Println("SMTP environment variables are not set correctly")
		return
	}

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		fmt.Println("Invalid SMTP port:", err)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", smtpEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	m.Attach(attachmentPath)

	d := gomail.NewDialer(smtpServer, port, smtpEmail, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send email:", err)
	}
}
