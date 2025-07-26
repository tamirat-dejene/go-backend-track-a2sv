# Use official Golang image
FROM golang:alpine3.22

# Set working directory
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the app
RUN go build -o task_manager_v1.2 .

# Run the app
CMD ["./task_manager_v1.2"]
