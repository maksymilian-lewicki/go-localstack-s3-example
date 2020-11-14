package example_test

import (
	example "go-localstack-s3-example"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS3PutObjectSuccessful(t *testing.T) {
	file, err := ioutil.ReadFile("images/go_logo.jpg")
	if err != nil {
		t.Error(err)
	}

	output, err := example.S3PutObject("go-localstack-s3-example.jpg", file)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, output.ETag)
}
