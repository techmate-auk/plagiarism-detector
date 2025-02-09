# Stage 1: Build
FROM golang:1.23.6-alpine3.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker's caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main main.go



# Stage 2: Runtime
FROM alpine:3.21

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy the templates directory (if your app uses it)
COPY --from=builder /app/templates ./templates

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD [ "/app/main" ]
