package models

type FromField struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Recipient struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type MailtrapEmail struct {
	From     FromField   `json:"from"`
	To       []Recipient `json:"to"`
	Subject  string      `json:"subject"`
	HTMLBody string      `json:"html"`
}
