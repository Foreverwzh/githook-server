
#build stage
FROM golang:alpine AS builder
COPY . /app
WORKDIR /app
ENTRYPOINT ["./main"]
EXPOSE 8080
