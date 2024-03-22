# Start from the official Golang image to make the build as straightforward as possible.
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o scm-api ./main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/scm-api .

# Command to run the executable.
CMD ["./scm-api"]
