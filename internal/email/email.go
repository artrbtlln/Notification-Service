package email

import (
	"fmt"
	"learn/internal/config"
	"learn/internal/model"
	"log"
	"net/smtp"
)

type Client struct {
	Config *config.Config
}

func NewClient(cfg *config.Config) *Client {
	return &Client{Config: cfg}
}
func (c *Client) Send(mail *model.Mail) error {
	addr := c.Config.MailHost + c.Config.MailPort

	auth := c.authsmtp()

	body := fmt.Sprintf("%s\n%s\n%v", mail.Header, mail.Body, mail.Attachments)

	if err := smtp.SendMail(addr, auth, c.Config.MailFrom, []string{mail.Email}, []byte(body)); err != nil {
		log.Println(err)
		return err
	}
	return nil

}
func (c *Client) authsmtp() smtp.Auth {
	auth := smtp.PlainAuth("", c.Config.MailFrom, c.Config.MailPassword, c.Config.MailHost)
	return auth
}
