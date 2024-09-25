package helpers

import (
	"email_service/models"
	"encoding/json"
	"fmt"
	"os"
)

func SendEmailVerification(data models.UserCredential) error {
	mailtrapUrl := os.Getenv("MAILTRAP_API_URL")
	authToken := os.Getenv("MAILTRAP_API_TOKEN")

	verificationUrl := os.Getenv("VERIFICATION_URL")
	url := fmt.Sprintf("%v/%v/%v", verificationUrl, data.ID, data.Token)

	htmlBody, err := VerificationEmailBody(data.Name, url)
	if err != nil {
		return err
	}

	payload := models.MailtrapEmailPayload{
		From: models.MailtrapEmailAddress{
			Email: "hello@iceiceice.biz.id",
			Name:  "FoxWash",
		},
		To: []models.MailtrapEmailAddress{
			{
				Email: data.Email,
			},
		},
		Subject:  "FoxWash Account Verification",
		Text:     "Please verify your email account.",
		HTML:     htmlBody,
		Category: "Email Verification",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", authToken),
		"Content-Type":  "application/json",
	}

	response, err := FetchAPI(mailtrapUrl, "POST", headers, payloadBytes)
	if err != nil {
		return err
	}

	fmt.Printf("email verification sent to %s\n with response: %v", data.Email, response)

	return nil
}

func SendEmailOrder(data models.Order) error {
	mailtrapUrl := os.Getenv("MAILTRAP_API_URL")
	authToken := os.Getenv("MAILTRAP_API_TOKEN")

	htmlBody, err := OrderEmailBody(data)
	if err != nil {
		return err
	}

	payload := models.MailtrapEmailPayload{
		From: models.MailtrapEmailAddress{
			Email: "hello@iceiceice.biz.id",
			Name:  "FoxWash",
		},
		To: []models.MailtrapEmailAddress{
			{
				Email: data.User.Email,
			},
		},
		Subject:  "Your Order Confirmation",
		Text:     "Thank you for your order!",
		HTML:     htmlBody,
		Category: "Order Confirmation",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", authToken),
		"Content-Type":  "application/json",
	}

	response, err := FetchAPI(mailtrapUrl, "POST", headers, payloadBytes)
	if err != nil {
		return err
	}

	fmt.Printf("order confirmation sent to %s\n with response: %v", data.User.Email, response)

	return nil
}

func SendPaymentSuccess(data models.PaymentSuccess) error {
	mailtrapUrl := os.Getenv("MAILTRAP_API_URL")
	authToken := os.Getenv("MAILTRAP_API_TOKEN")

	htmlBody, err := PaymentSuccessEmailBody(data)
	if err != nil {
		return err
	}

	payload := models.MailtrapEmailPayload{
		From: models.MailtrapEmailAddress{
			Email: "hello@iceiceice.biz.id",
			Name:  "FoxWash",
		},
		To: []models.MailtrapEmailAddress{
			{
				Email: data.Email,
			},
		},
		Subject:  "Your Payment Successfully!",
		Text:     "Washer is preparing🛵",
		HTML:     htmlBody,
		Category: "Payment Successful",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", authToken),
		"Content-Type":  "application/json",
	}

	response, err := FetchAPI(mailtrapUrl, "POST", headers, payloadBytes)
	if err != nil {
		return err
	}

	fmt.Printf("payment successful sent to %s\n with response: %v", data.Email, response)

	return nil
}
