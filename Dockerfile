FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./app/cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY app/config/config.yaml /config/config.yaml


EXPOSE 8080

CMD ["./main"]
