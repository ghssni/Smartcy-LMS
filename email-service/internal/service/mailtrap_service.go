package service

//func SendMailtrapEmail(email, subject, htmlBody string) error {
//	apiKey := os.Getenv("MAILTRAP_API_KEY")
//	url := os.Getenv("MAILTRAP_DOMAIN")
//
//	to := []models.Recipient{
//		{
//			Email: email,
//			Name:  "Confirm Payment",
//		},
//	}
//	from := models.FromField{
//		Email: "no-reply@mailtrap.io",
//		Name:  "Smartcy LMS",
//	}
//	emailData := models.MailtrapEmail{
//		From:     from,
//		To:       to,
//		Subject:  subject,
//		HTMLBody: htmlBody,
//	}
//
//	jsonData, err := json.Marshal(emailData)
//	if err != nil {
//		return fmt.Errorf("error marshaling JSON: %v", err)
//	}
//
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
//	if err != nil {
//		return fmt.Errorf("error creating request: %v", err)
//	}
//
//	req.Header.Set("Authorization", "Bearer "+apiKey)
//	req.Header.Set("Content-Type", "application/json")
//
//	client := &http.Client{Timeout: time.Second * 10}
//	resp, err := client.Do(req)
//
//	if err != nil {
//		return fmt.Errorf("error sending request: %v", err)
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode >= 400 {
//		bodyBytes, _ := ioutil.ReadAll(resp.Body)
//		bodyString := string(bodyBytes)
//		return fmt.Errorf("error: received non-2xx response code: %d, body: %s", resp.StatusCode, bodyString)
//	}
//	return nil
//}
//
//func SendEmailPayment(email, courseName, paymentURL string) error {
//	subject := "Payment Confirmation for Course"
//	htmlBody := fmt.Sprintf(`
//		<html>
//			<body>
//				<h2>Payment Confirmation</h2>
//				<p>Dear Student,</p>
//				<p>Your payment for the course <b>%s</b> is pending.</p>
//				<p>Please complete the payment using the following link:</p>
//				<a href="%s">Click here to pay</a>
//				<p>Thank you!</p>
//			</body>
//		</html>`, courseName, paymentURL)
//
//	return SendMailtrapEmail(email, subject, htmlBody)
//}
//
//func SendEmailSuccess(email, courseName string) error {
//	subject := "Payment Confirmation for Course"
//	htmlBody := fmt.Sprintf(`
//		<html>
//			<body>
//				<h2>Payment Confirmation</h2>
//				<p>Dear Student,</p>
//				<p>Your payment for the course <b>%s</b> has been successfully processed.</p>
//				<p>Thank you for your payment!</p>
//			</body>
//		</html>`, courseName)
//
//	return SendMailtrapEmail(email, subject, htmlBody)
//}
//
//func SendEmailForgotPassword(email, resetURL, resetToken string) error {
//	subject := "Reset Password Request"
//	fullResetURL := fmt.Sprintf("%s?token=%s", resetURL, resetToken)
//	htmlBody := fmt.Sprintf(`
//		<html>
//			<body>
//				<h2>Reset Password</h2>
//				<p>Dear Student,</p>
//				<p>We have received a request to reset your password.</p>
//				<p>Please click the link below to reset your password:</p>
//				<a href="%s">Reset Password</a>
//				<p>If you did not request this, please ignore this email.</p>
//				<p>Thank you!</p>
//			</body>
//		</html>`, fullResetURL)
//
//	return SendMailtrapEmail(email, subject, htmlBody)
//}
