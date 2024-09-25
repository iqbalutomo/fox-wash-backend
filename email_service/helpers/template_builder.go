package helpers

import (
	"bytes"
	"email_service/models"
	"fmt"
	"text/template"
)

func VerificationEmailBody(name, url string) (string, error) {
	template, err := template.ParseFiles("./templates/email_verification.html")
	if err != nil {
		return "", fmt.Errorf("failed to parsing file: %v", err)
	}

	templateData := map[string]string{
		"Name": name,
		"URL":  url,
	}

	buf := new(bytes.Buffer)
	if err := template.Execute(buf, templateData); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func OrderEmailBody(data models.Order) (string, error) {
	template, err := template.ParseFiles("./templates/email_order_user.html")
	if err != nil {
		return "", fmt.Errorf("failed to parsing file: %v", err)
	}

	templateData := map[string]interface{}{
		"order_detail": data.OrderDetail,
		"User":         data.User,
		"Washer":       data.Washer,
		"Address":      data.Address,
		"Payment":      data.Payment,
		"Status":       data.Status,
	}

	buf := new(bytes.Buffer)
	if err := template.Execute(buf, templateData); err != nil {
		return "", err
	}

	return buf.String(), nil
}
