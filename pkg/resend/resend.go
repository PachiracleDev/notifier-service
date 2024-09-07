package resend

import (
	"notifier/config"

	"github.com/resend/resend-go/v2"
)

type ResendImplementation struct {
	client *resend.Client
	from   string
}

func NewResendImplementation(conf *config.Config) *ResendImplementation {
	client := resend.NewClient(conf.Resend.ApiKey)
	return &ResendImplementation{
		client: client,
		from:   conf.Resend.From,
	}
}

type SendEmailResendDto struct {
	Recipient string
	Subject   string
	HtmlBody  string
}

func (u *ResendImplementation) SendEmail(dto SendEmailResendDto) error {
	params := &resend.SendEmailRequest{
		To:      []string{dto.Recipient},
		From:    u.from,
		Subject: dto.Subject,
		Html:    dto.HtmlBody,
	}

	_, err := u.client.Emails.Send(params)
	if err != nil {
		panic(err)
	}

	return nil
}
