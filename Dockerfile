FROM golang:1.23-alpine AS builder

# Install git
RUN apk add --no-cache git

# Set working directory to root of the container
WORKDIR /app

# Copy the ENTIRE project into the container
# (This ensures we have all subfolders: backend/auth-service, backend/webhook-service, etc.)
COPY . .

# Accept the target service path as a build argument
# Example: backend/webhook-service
ARG SERVICE_PATH

# Change directory to the specific service folder
WORKDIR /app/${SERVICE_PATH}

# Download dependencies for THAT specific service
RUN go mod download

# Build the binary
RUN go build -o /server main.go

# --- Final Stage ---
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /server .

# Expose port (just for documentation)
EXPOSE 8080

# Run the server
CMD ["./server"]