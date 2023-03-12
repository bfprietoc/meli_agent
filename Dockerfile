# Set the base image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY go.mod ./
COPY go.sum ./

# Build the Go application
RUN go build -o meli

# Expose the port the application will run on
EXPOSE 8080

# Set the entry point for the container to start the application
ENTRYPOINT ["./meli"]