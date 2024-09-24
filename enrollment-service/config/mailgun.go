package config

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// SendEmailPayment sends an email to the student with a payment URL
func SendEmailPayment(email, courseName, paymentURL string) error {
	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")

	mg := mailgun.NewMailgun(domain, apiKey)

	sender := fmt.Sprintf("Smartcy LMS <no-reply@%s>", domain)
	subject := "Payment Confirmation for Course"
	htmlBody := fmt.Sprintf(`
		<html>
			<body>
				<h2>Payment Confirmation</h2>
				<p>Dear Student,</p>
				<p>Your payment for the course <b>%s</b> is pending.</p>
				<p>Please complete the payment using the following link:</p>
				<a href="%s">Click here to pay</a>
				<p>Thank you!</p>
			</body>
		</html>`, courseName, paymentURL)
	recipient := email

	message := mg.NewMessage(sender, subject, "", recipient)

	message.SetHtml(htmlBody)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mg.Send(ctx, message)
	if err != nil {
		logrus.Println("Error sending email:", err)
	}
	return nil
}

// SendEmailSuccess sends an email to the student confirming the payment
func SendEmailSuccess(email, courseName string) error {
	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")

	mg := mailgun.NewMailgun(domain, apiKey)

	sender := fmt.Sprintf("Smartcy LMS <no-reply@%s>", domain)
	subject := "Payment Confirmation for Course"
	htmlBody := fmt.Sprintf(`
		<html>
			<body>
				<h2>Payment Confirmation</h2>
				<p>Dear Student,</p>
				<p>Your payment for the course <b>%s</b> has been successfully processed.</p>
				<p>Thank you for your payment!</p>
			</body>
		</html>`, courseName)
	recipient := email

	message := mg.NewMessage(sender, subject, "", recipient)

	message.SetHtml(htmlBody)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mg.Send(ctx, message)
	if err != nil {
		logrus.Println("Error sending email:", err)
	}
	return nil
}
