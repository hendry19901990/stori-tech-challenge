# Use the official Go image as the base image
FROM golang:1.17-alpine as builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o app

# Final stage: create a lightweight image
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder app .

# Set executable permissions for the app binary
RUN chmod +x app

# Start the application
CMD ["./app"]