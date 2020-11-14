FROM golang:1.15.5-alpine3.12
ENV GO111MODULE="on"
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build binary
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build ./...
