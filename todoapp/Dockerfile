# Stage 1: Build the Go application
FROM golang:1.23-alpine as builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY src/go.mod src/go.sum ./
RUN go mod download

# Copy the entire application source code
COPY src/ ./
RUN go build -o todoapp main.go

# Stage 2: Create a minimal runtime container
FROM alpine:latest

WORKDIR /app

# Copy the built binary and static assets from the builder stage
COPY --from=builder /app/todoapp .
COPY src/html ./html
COPY src/images ./images

EXPOSE 8080

CMD ["./todoapp"]

