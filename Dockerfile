# Use an official Golang runtime as a parent image
FROM golang:1.18-alpine3.15

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Install any needed packages
RUN apk add --no-cache git

# Fetch dependencies
RUN go mod download

# Build the Go app
RUN go build -o /go/bin/app

# Set the entry point to the app binary
ENTRYPOINT ["/go/bin/app"]
