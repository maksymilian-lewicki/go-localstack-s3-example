# go localstack s3 example
[![Actions Status](https://github.com/maksymilian-lewicki/go-localstack-s3-example/workflows/build-test/badge.svg)](https://github.com/maksymilian-lewicki/go-localstack-s3-example/actions)
[![Coverage Status](https://coveralls.io/repos/github/maksymilian-lewicki/go-localstack-s3-example/badge.svg)](https://coveralls.io/github/maksymilian-lewicki/go-localstack-s3-example)
[![Go Report Card](https://goreportcard.com/badge/github.com/maksymilian-lewicki/go-localstack-s3-example)](https://goreportcard.com/report/github.com/maksymilian-lewicki/go-localstack-s3-example) 

Example of a simple application in Golang, which shows how to use S3 from localstack.

## localstack in docker
[docker-compose.yaml localstack](./docker-compose.yml#L11-L21) - defined `localstack` service for s3, exposed on port `4566`. `localstack-data` volume used to persist data after rebuild.

[docker-compose.yaml wait-for-localstack](./docker-compose.yml#L23-L27) - service used to check if the localstack service is ready.

[test](scripts/test.sh) - script to run tests. The script creates a bucket in s3 in a line number [7](scripts/test.sh#L7).

[s3.go](./s3.go#L13-L38) - code responsible for uploading files one s3.