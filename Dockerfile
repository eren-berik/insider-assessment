FROM golang:1.23.6-alpine AS builder

WORKDIR /app

# Copy the entire project first
COPY . .

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger documentation
RUN swag init -g pkg/api/server.go

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o messaging-app ./cmd/messaging-app

# Use a minimal alpine image for the final stage
FROM alpine:latest

WORKDIR /app

# Copy only the necessary files from builder
COPY --from=builder /app/messaging-app .
COPY --from=builder /app/docs ./docs

EXPOSE 4300

CMD ["./messaging-app"]