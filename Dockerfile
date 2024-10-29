# Stage 1: Build the Go binary
FROM golang:1.23.2 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp .

# Stage 2: Run the Go binary in a minimal environment
FROM alpine:3.18

# Install necessary certificates for making HTTPS requests
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled Go binary from the build stage
COPY --from=build /app/myapp .

# Ensure the binary has executable permissions
RUN chmod +x ./myapp

# Expose the port that the application runs on
EXPOSE 8080

# Command to run the Go application
CMD ["./myapp"]