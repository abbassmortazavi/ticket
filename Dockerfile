FROM golang:1.25.3

LABEL authors="abbass"

WORKDIR /app

# Install air for live reload
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application with explicit architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Add execute permissions
RUN chmod +x main

# Expose port
EXPOSE 8080

# Run the application
#CMD ["./main"]
CMD ["air", "-c", ".air.toml"]