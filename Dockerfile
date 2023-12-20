FROM golang:1.21.5-bullseye

WORKDIR /app

COPY . .
RUN go mod tidy
