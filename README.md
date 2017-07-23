# gcs-api
This is a toy API for GCS. It serves as an example of how an API can retrieve
information from a database and serve it back in a useable JSON format.

# How to build?

This is a Go application, so it may require you to `go get` the application
dependencies that are needed for the application itself.

Another dependency of this application is that a MySQL database is needed. See
the `testdata.sql` as a starting point for setting it up. Also, take a look at
the database helper.

# How to run?

Once you have all the dependencies that you need, then it should be just a:

    go run main.go

# How to run with Docker?

First you need to build the image:

    docker build -t gcsapi .

Once that has been done we can run it:

    docker run -d -p 8080:8080 gcsapi gcsapi

# Licensing

  MIT.
