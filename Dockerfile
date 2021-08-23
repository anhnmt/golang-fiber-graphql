# ./Dockerfile

FROM golang:alpine as builder

# Add Maintainer Info
LABEL maintainer="Tuan Anh Nguyen Manh <xdorro@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/config.yml .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
ENTRYPOINT ["./main"]