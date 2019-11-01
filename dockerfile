FROM golang:1.13
COPY . /app
WORKDIR /app
RUN go get github.com/gin-gonic/gin
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]