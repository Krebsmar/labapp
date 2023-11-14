# Start from the official Go image.
FROM golang:1.21.4-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code into the container.
COPY . .

# Build the application.
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Update openssl/libcrypto3
RUN apk --no-cache upgrade openssl/libcrypto3

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose the port on which the application runs.
EXPOSE 8080

# Set the command to run the application.
CMD ["./main"]

# Add a healthcheck
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD wget --quiet --tries=1 --spider http://localhost:8080 || exit 1