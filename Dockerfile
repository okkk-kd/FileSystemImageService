FROM golang:1.19.5

WORKDIR /usr/src/app

COPY .. .
RUN go mod tidy