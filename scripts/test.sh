#!/usr/bin/env bash
set -e

docker-compose build --pull app
docker-compose up -d localstack
docker-compose run --rm wait-for-localstack
docker-compose exec localstack awslocal s3 mb s3://go-localstack-s3-example
docker-compose run --rm app go test -v -covermode=count ./...
docker-compose down -v
