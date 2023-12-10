package model

type Mail struct {
	Header      string `json:"header"` // РУСТАМА В рот любил ЭТО МОДЕЛЬ
	Body        string `json:"body"`
	Email       string `json:"email"`
	Attachments string `json:"attachments,omitempty"`
}
