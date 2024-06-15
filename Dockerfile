# Use the official Go image as the base image
FROM golang:1.22.4-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY . .

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o bin ./cmd/web

# Expose the port that the application runs on
EXPOSE 8080

# Start the application
CMD [ "bin/web" ]