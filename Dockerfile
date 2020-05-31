FROM golang:1.14.3

WORKDIR /app

ADD . /app

RUN go build server.go

CMD ["./server"]