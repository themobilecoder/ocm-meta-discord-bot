# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src/*.go ./
RUN go build -o /docker-ocm-meta-bot

CMD ["/docker-ocm-meta-bot"]