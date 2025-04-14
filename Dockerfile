# Development image based on Go
FROM golang:1.24-alpine

# Install air for hot reloading
RUN go install github.com/air-verse/air@latest

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Expose port
EXPOSE 8080

# Run air for hot reloading
CMD ["air"]