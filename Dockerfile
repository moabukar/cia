# Build stage
FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

# Build the Go app
RUN go build -o cia

### Release stage
FROM alpine:latest

# Set the Current Working Directory
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/cia .

# Command to run the executable
CMD ["./cia"]
