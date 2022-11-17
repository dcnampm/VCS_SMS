#from golang base image //FROM <image>[:<tag>] [AS <name>]
FROM golang:alpine as builder

#install Git
RUN apk update && apk add --no-cache git

#Working directory in container
WORKDIR /app

COPY go.mod go.sum /app/

#dowload dependencies
RUN go mod download 

COPY . /app/

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the app.env file
COPY --from=builder /app/main .
COPY --from=builder /app/app.env .       

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]