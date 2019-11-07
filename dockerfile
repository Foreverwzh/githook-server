
#build stage
FROM golang:alpine AS builder

ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io

RUN apk update && apk add git openssh
COPY . /app
WORKDIR /app
RUN mkdir /root/.ssh/
RUN cp .ssh/id_rsa /root/.ssh/id_rsa

RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN go build main.go
ENTRYPOINT ["./main"]
EXPOSE 8080
