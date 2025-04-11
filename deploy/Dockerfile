# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/build/sonda ./cmd/sonda

FROM alpine:3.18

# Add ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Create a non-root user and group
RUN addgroup -S sonda && adduser -S sonda_worker -G sonda

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/build/sonda .
# Copy configuration
COPY configs/config.yaml /app/configs/

RUN chown -R sonda_worker:sonda /app

USER sonda_worker

EXPOSE 8080

CMD ["/app/sonda"]
