
#build stage
FROM golang:alpine AS builder

ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io

RUN apk update && apk add git
COPY . /app
WORKDIR /app
RUN go build main.go
ENTRYPOINT ["./main"]
EXPOSE 8080
