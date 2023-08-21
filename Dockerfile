# Use the official Go image as the base image
FROM golang:1.20.7-alpine3.18

RUN apk update && apk add --no-cache git
# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY main.go .

# Build the Go binary
RUN go build main.go  

# Specify the command to run when the container starts
CMD ["./main"]
