FROM golang:latest

ENV DATABASE_HOST 192.168.99.100:3306
ENV DATABASE_NAME gcstest
ENV DATABASE_PROTOCOL tcp
ENV DATABASE_USERNAME root
ENV DATABASE_PASSWORD secret

COPY . /go/src/github.com/mickelsonm/gcs-api

WORKDIR /go/src/github.com/mickelsonm/gcs-api

RUN go get && go build -o API ./main.go

ENTRYPOINT /go/src/github.com/mickelsonm/gcs-api/API

EXPOSE 8080
