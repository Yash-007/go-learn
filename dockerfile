# Use official Golang image to build the app
FROM golang AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build the Go app
RUN go build -o myapp

# Use a minimal base image to run the app
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/myapp .

# Expose a port (if your app runs on a specific port)
EXPOSE 8080

# Run the app
CMD ["./myapp"]
