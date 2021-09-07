FROM golang:1.16-alpine  AS builder

RUN apk add --update make

WORKDIR /app

COPY . .

RUN go mod download
