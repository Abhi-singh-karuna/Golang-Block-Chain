# Step 1: Use an official Go image to build the application
FROM golang:1.20-alpine AS build

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Step 4: Download dependencies
RUN go mod download

# Step 5: Copy the rest of the application source code
COPY . .

# Step 6: Build the Go server (binary named 'server')
RUN go build -o server main.go mycc.go

# Step 7: Use a smaller base image for the final container
FROM alpine:latest

# Step 8: Set the working directory for the final container
WORKDIR /app

# Step 9: Copy the Go server binary from the builder container
COPY --from=build /app/server .

# Step 10: Expose the port that the Gin server will run on
EXPOSE 8080

# Step 11: Start the Go server
CMD ["./server"]
