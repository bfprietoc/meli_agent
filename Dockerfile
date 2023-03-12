# Set the base image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .


# Expose the port the application will run on
EXPOSE 8080

# Set the entry point for the container to start the application
ENTRYPOINT ["./meli"]