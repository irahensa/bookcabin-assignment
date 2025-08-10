FROM golang:1.24-alpine

RUN apk add --no-cache musl-dev gcc

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1
RUN go build -o main cmd/main.go

EXPOSE 8080

CMD ["./main"]