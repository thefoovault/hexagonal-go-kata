# Use the official Golang image as the base image
FROM golang:1.22-alpine as builder

# Set the working directory inside the container
WORKDIR /app

#COPY ./src .
COPY . .

# Initialize a new Go module if go.mod does not exist
RUN if [ ! -f go.mod ]; then go mod init go_test; fi

# Download all dependencies
RUN go mod tidy

# Build the Go application
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/api/main.go

# Use a lightweight base image for the final image
FROM alpine:latest

# Install ca-certificates
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
#COPY --from=builder /app/main .
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run the Go application
#CMD ["./main"]
ENTRYPOINT ["./main"]
