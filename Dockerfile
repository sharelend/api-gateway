# Use official Go image
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy all files into the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the application
RUN go build -o api-gateway .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./api-gateway"]
