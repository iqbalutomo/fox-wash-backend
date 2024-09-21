package models

type MailtrapEmailPayload struct {
	From     MailtrapEmailAddress   `json:"from"`
	To       []MailtrapEmailAddress `json:"to"`
	Subject  string                 `json:"subject"`
	Text     string                 `json:"text"`
	HTML     string                 `json:"html"`
	Category string                 `json:"category,omitempty"`
}

type MailtrapEmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}
