FROM golang:latest

WORKDIR /app
COPY . /app

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  GO111MODULE=on
EXPOSE 1234


CMD ["go", "run", "main.go"]
