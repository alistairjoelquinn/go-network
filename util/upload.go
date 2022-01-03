package util

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadImage(filename string) (string, error) {
	imageFilePath := fmt.Sprintf("%s/%s", Env("UPLOAD_PATH"), filename)

	f, err := os.Open(imageFilePath)
	if err != nil {
		return "", err
	}

	s3Session := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(Env("AWS_KEY"), Env("AWS_SECRET"), ""),
	}))

	uploader := s3manager.NewUploader(s3Session)
	aclValue := "public-read"

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(Env("AWS_BUCKET")),
		Key:    aws.String(filename),
		Body:   f,
		ACL:    &aclValue,
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
