# Use Golang base image
FROM golang:1.24

# Set the working directory
WORKDIR /app

# Copy project files
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o orderservice main.go

# Expose port
EXPOSE 8000

# Run the application
CMD ["/app/orderservice"]