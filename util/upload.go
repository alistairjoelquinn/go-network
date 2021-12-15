package util

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadImage(filename string) error {
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)

	imageFilePath := fmt.Sprintf("%s/%s", Env("UPLOAD_PATH"), filename)

	f, err := os.Open(imageFilePath)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	results, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(Env("AWS_BUCKET")),
		Key:    aws.String(Env("AWS_KEY")),
		Body:   f,
	})
	log.Println("results", results, err)

	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", filename)

	return nil
}
