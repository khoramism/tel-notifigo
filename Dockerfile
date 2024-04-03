# syntax=docker/dockerfile:1

FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-notification-application

EXPOSE 8787

# Run

CMD ["/golang-notification-application"]