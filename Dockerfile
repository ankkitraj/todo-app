# Use the official Golang image to create a build artifact.
FROM golang:1.22.2 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files do not change
RUN go mod download

# Copy the entire source directory contents into the container
COPY . .

# Change to the directory containing Go files
WORKDIR /app/src

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ../todo-app

# Use Google's Distroless base image for the final image
FROM gcr.io/distroless/base-debian10

# Copy the built binary from the builder stage
COPY --from=builder /app/todo-app /

# Command to run
CMD ["/todo-app"]

