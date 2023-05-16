# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="sauravsuman2920@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. 
# Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./src ./src

# Build the Go app
RUN go build -o main ./src/cmd

# Expose port 8000 to the outside world
EXPOSE 8000

# Run the executable
CMD ["./main"]
