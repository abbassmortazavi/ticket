FROM golang:1.24.6

LABEL authors="abbass"

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main ./cmd/api

# Add execute permissions
RUN chmod +x main

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]