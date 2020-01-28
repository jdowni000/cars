FROM golang:alpine

ADD . /app

WORKDIR /app/cmd/car-server

RUN go build -v

ENTRYPOINT /app/cmd/car-server/car-server

