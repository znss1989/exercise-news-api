# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY docs/ ./docs/

RUN go mod download

COPY *.go ./

RUN apk add build-base
RUN go get .

EXPOSE 8080
CMD [ "go", "run", "." ]

# RUN go build -o /mm-news-server

# CMD [ "/mm-news-server" ]