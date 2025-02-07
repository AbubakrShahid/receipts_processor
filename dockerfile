# Use the official Go image with version 1.23.6
FROM golang:1.23.6 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod tidy

# Copy the entire source code to /app inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from an Ubuntu image for newer glibc
FROM ubuntu:latest

# Install dependencies, including libc6
RUN apt-get update && apt-get install -y \
    libc6 \
    libc6-dev \
    libstdc++6 \
    ca-certificates \
    curl

# Set the Current Working Directory inside the container for the second stage
WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 for the app
EXPOSE 8080

# Run the Go application
CMD ["./main"]
