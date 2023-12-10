package service

import (
	"learn/internal/email"
	"learn/internal/model"
	"log"
)

type EmailSrv struct {
	Client *email.Client
}

func NewEmailSrv(client *email.Client) *EmailSrv {
	return &EmailSrv{Client: client}
}
func (s *EmailSrv) Send(mail *model.Mail) error {
	if err := s.Client.Send(mail); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
