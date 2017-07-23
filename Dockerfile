FROM scratch

#You may need to change these environment variables depending on your setup
ENV DATABASE_HOST 192.168.99.100:3306
ENV DATABASE_NAME gcstest
ENV DATABASE_PROTOCOL tcp
ENV DATABASE_USERNAME root
ENV DATABASE_PASSWORD secret

#first we need to build it like so:
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gcsapi .

ADD gcsapi /
EXPOSE 8080
ENTRYPOINT ["/gcsapi"]
