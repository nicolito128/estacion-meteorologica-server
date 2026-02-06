# Imagen base
FROM golang:1.25-alpine3.22 AS base

ENV DIR /server

WORKDIR $DIR

COPY . .

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]
