FROM golang:latest

#You may need to change these environment variables depending on your setup
ENV DATABASE_HOST 192.168.99.100:3306
ENV DATABASE_NAME gcstest
ENV DATABASE_PROTOCOL tcp
ENV DATABASE_USERNAME root
ENV DATABASE_PASSWORD secret

#copies all contents into the src project directory
COPY . /go/src/github.com/mickelsonm/gcs-api

#sets our current work directory inside of the src project directory
WORKDIR /go/src/github.com/mickelsonm/gcs-api

#gets our dependencies and builds the API binary file
RUN go get && go build -o API ./main.go

#sets the entrypoint for this application to be the API
ENTRYPOINT /go/src/github.com/mickelsonm/gcs-api/API

#exposes 8080, which is the port the API listens on
EXPOSE 8080
