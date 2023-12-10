package service

import (
	"learn/internal/email"
	"learn/internal/model"
)

type Email interface {
	Send(mail *model.Mail) error
}

type Serivce struct {
	Email
	Client *email.Client
}

func NewSerivce(c *email.Client) *Serivce {
	return &Serivce{
		Email:  NewEmailSrv(c),
		Client: c,
	}
}
