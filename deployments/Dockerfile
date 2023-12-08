# Step 1: Use an official Go runtime as a parent image
FROM golang:1.21-alpine as builder

# Step 2: Set the working directory inside the container
WORKDIR /go/src/app

# Step 3: Copy the local package files to the container's workspace.
COPY . .

# Step 4: Download any necessary dependencies.
# This step can be optimized with Go modules.
RUN go mod download

# Step 5: Build the application for a specific platform.
# Here, we build for a 64-bit Linux system.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app

# Step 6: Use a minimal alpine image for the production container
FROM alpine:latest

# Step 7: Copy the binary from the builder stage to the production container
COPY --from=builder /go/bin/app /go/bin/app

# Step 8: Set the binary as the entrypoint of the container
ENTRYPOINT ["/go/bin/app"]

# Optional: Expose ports. Only necessary if your application is a server.
# EXPOSE 8080