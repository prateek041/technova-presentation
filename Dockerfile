# Base image
FROM golang:1.19

# Decide the working directory
WORKDIR /app

# Copy the go mode file
COPY go.mod ./

# Download the dependencies
RUN go mod download

# Copy the remaining code
COPY *.go ./

# Build and Run the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /hello-app
CMD ["/hello-app"]