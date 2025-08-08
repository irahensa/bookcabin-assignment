FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY vouchers.db .

RUN go mod download

COPY . .

RUN  go build -o main cmd/main.go

EXPOSE 8080

CMD ["./main"]