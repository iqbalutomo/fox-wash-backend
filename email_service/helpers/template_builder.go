package helpers

import (
	"bytes"
	"fmt"
	"text/template"
)

func VerificationEmailBody(name, url string) (string, error) {
	template, err := template.ParseFiles("./templates/email_verification.html")
	if err != nil {
		return "", fmt.Errorf("failed to parsing file: %v", err)
	}

	templateData := map[string]string{
		"Logo": "https://raw.githubusercontent.com/iqbalutomo/fox-wash-backend/refs/heads/master/assets/foxwash-logo.png",
		"Name": name,
		"URL":  url,
	}

	buf := new(bytes.Buffer)
	if err := template.Execute(buf, templateData); err != nil {
		return "", err
	}

	return buf.String(), nil
}
