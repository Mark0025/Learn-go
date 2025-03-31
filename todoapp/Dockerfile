# Stage 1: Build the Go application
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY src/go.mod src/go.sum ./
RUN go mod download

# Copy the entire application source code
COPY src/ ./

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todoapp main.go

# Stage 2: Create a minimal runtime container
FROM alpine:latest

WORKDIR /app

# Copy the built binary and static assets from the builder stage
COPY --from=builder /app/todoapp .
COPY src/html ./html
COPY src/images ./images

# Add non-root user for security
RUN adduser -D -g '' appuser
USER appuser

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./todoapp"]

