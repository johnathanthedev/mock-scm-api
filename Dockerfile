# Start from the official Golang image to make the build as straightforward as possible.
FROM golang:1.18 as builder

# Create a directory for your app.
WORKDIR /app

# Copy the go.mod and go.sum file into the directory.
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the rest of your application's source code into the directory.
COPY . .

# Build your application.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o scm-api ./main.go

# Start a new stage from scratch.
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/scm-api .

# Command to run the executable.
CMD ["./scm-api"]
