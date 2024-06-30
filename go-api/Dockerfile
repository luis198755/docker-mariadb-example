# Use the official Golang image to build the app
FROM golang:alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY ./main.go ./main.go

# Download any dependencies
RUN go mod init example.com/main

# Download any dependencies
RUN go mod tidy

# Compile the program
RUN go build -o  main

# Expose port 8080 to the outside world
EXPOSE 80

# Step 7: Command to run the executable
CMD ["./main"]