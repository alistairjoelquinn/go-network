package util

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendResetEmail(code string, email string) error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region:      aws.String("eu-central-1"),
			Credentials: credentials.NewStaticCredentials(Env("AWS_KEY"), Env("AWS_SECRET"), ""),
		},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(Env("EMAIL")),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(fmt.Sprintf("You have just made a request to change your password, please enter the following code on the commonground reset page in order to proceed: \n\n %s \n\n\n Please do not reply to this email.", code)),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("Password Reset Code"),
			},
		},
		Source: aws.String(Env("EMAIL")),
	}

	_, err = svc.SendEmail(input)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
