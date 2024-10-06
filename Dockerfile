# Development Stage
# FROM golang:1.23-alpine AS development

# # Install air for hot-reloading
# RUN go install github.com/air-verse/air@latest

# # Set working directory
# WORKDIR /app

# # Copy go.mod and go.sum
# COPY go.mod go.sum ./

# # Download Go dependencies
# RUN go mod download

# # Copy the rest of the application code
# COPY . .

# # Expose the port for development
# EXPOSE 8080

# # Command to run air for development (hot-reload)
# CMD ["air", "-c", ".air.toml"]

# Production Stage
FROM golang:1.23-alpine AS development

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/api

# Expose the port for production
EXPOSE 8080

# Command to run the application in production
CMD ["./main"]
