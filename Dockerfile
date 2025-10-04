# ----------------------------------------------------------------------
# STAGE 1: Build the Go application
# ----------------------------------------------------------------------
FROM golang:1.25.1 AS builder

# Set necessary environment variables for Go modules
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to cache dependencies
COPY go.mod .
COPY go.sum .

# Download all dependencies. Dependencies will be cached unless go.mod/go.sum change
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# -o /dist/main tells Go to place the compiled binary in /dist/main
# ./... tells Go to compile all packages in the current directory and its subdirectories
RUN go build -ldflags "-s -w" -o /dist/main ./...


# ----------------------------------------------------------------------
# STAGE 2: Create the final, minimal image
# ----------------------------------------------------------------------
FROM alpine:latest AS final

# Optional: Install CA certificates if your application needs to make HTTPS requests
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the 'builder' stage
# The path /dist/main must match the path used in the RUN go build command
COPY --from=builder /dist/main .

# Expose the port your Gin application listens on (default is 8080)
EXPOSE 8080

# The command to run the executable
CMD ["./main"]