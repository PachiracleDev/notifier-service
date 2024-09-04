package awssdk

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type SDKImplementation struct {
	svc *ses.SES
}

func NewSDKImplementation() *SDKImplementation {
	sess, err := session.NewSessionWithOptions(session.Options{
		//Default profile, read from ~/.aws/credentials and ~/.aws/config
		Config: aws.Config{
			Region: aws.String("sa-east-1"),
		},
	})

	if err != nil {
		log.Fatalf("Error creando la sesión de AWS: %v", err)
	}

	svc := ses.New(sess)

	return &SDKImplementation{
		svc: svc,
	}
}

type SendEmailDto struct {
	Recipient string
	Subject   string
	HtmlBody  string
}

func (s *SDKImplementation) SendEmail(dto SendEmailDto) error {

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(dto.Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(dto.HtmlBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(dto.Subject),
			},
		},
		Source: aws.String(os.Getenv("AWS_EMAIL")),
	}

	_, err := s.svc.SendEmail(input)

	if err != nil {
		fmt.Println("Error sending email: ", err)
		return err
	}

	return nil
}
