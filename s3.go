package example

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3PutObject Upload file to S3
func S3PutObject(path string, file []byte) (output *s3.PutObjectOutput, err error) {
	awsEndpoint := os.Getenv("AWS_ENDPOINT")
	awsRegion := os.Getenv("AWS_REGION")
	s3Bucket := os.Getenv("S3_BUCKET")

	s := session.Must(session.NewSession(&aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(awsEndpoint),
		Region:           aws.String(awsRegion),
	}))

	c := s3.New(s, &aws.Config{})

	fileContentType := http.DetectContentType(file)
	reader := bytes.NewReader(file)

	input := &s3.PutObjectInput{
		Bucket:        aws.String(s3Bucket),
		Key:           aws.String(path),
		Body:          reader,
		ContentType:   aws.String(fileContentType),
		ContentLength: aws.Int64(reader.Size()),
	}

	return c.PutObject(input)
}
