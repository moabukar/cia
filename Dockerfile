# Build stage
FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

# Build the Go app
RUN go build -o cia

### Release stage
FROM alpine:latest

# Set the Current Working Directory
WORKDIR /

COPY --from=builder /app/cia .

# nobody user
USER 65534 

CMD ["./cia"]
