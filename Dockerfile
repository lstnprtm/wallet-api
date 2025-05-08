# Start from official Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Install git and MySQL client
RUN apk add --no-cache git mysql-client

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the source
COPY wallet-api-final .

# Build the app
RUN go build -o wallet-api ./cmd/server

# Expose port
EXPOSE 8080

# Run
CMD ["./wallet-api"]
