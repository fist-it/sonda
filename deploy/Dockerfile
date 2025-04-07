FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o sonda cmd/sonda/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/sonda .

EXPOSE 8080

CMD ["./sonda"]

