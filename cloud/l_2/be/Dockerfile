# Stage 1: Build the Go application
FROM golang:1.21 AS builder

WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application code and build the binary for Linux
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Minimal runtime image
FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Ensure the binary has executable permissions
RUN chmod +x main

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
