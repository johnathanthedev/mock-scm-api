FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o scm-api ./main.go
RUN CGO_ENABLED=0 GOOS=linux go build -v -o migrate ./cmd/cli-tools/migrate/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -v -o rollback ./cmd/cli-tools/rollback/main.go

COPY db/migrations /app/db/migrations

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/scm-api .
COPY --from=builder /app/migrate .
COPY --from=builder /app/rollback .
COPY --from=builder /app/db /app/db

CMD ["./scm-api"]
