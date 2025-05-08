# Use Go 1.24 release candidate based on Debian Bullseye
FROM golang:1.24-bullseye

# Install git (needed for some go packages)
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the Go app
RUN go build -o server .

# Expose the port your app uses
EXPOSE 8080

# Run the app
CMD ["./server"]
